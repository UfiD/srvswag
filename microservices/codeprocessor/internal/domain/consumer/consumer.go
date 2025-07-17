package consumer

import (
	"bytes"
	"codeproc/microservices/codeprocessor/internal/domain/consumer/types"
	domain "codeproc/microservices/codeprocessor/internal/domain/entity"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types/build"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/moby/go-archive"
)

type Consumer struct{}

func NewConsumer() *Consumer {
	return &Consumer{}
}

func (c *Consumer) Do(object domain.Object) *domain.ObjectResult {
	exeDir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("не удалось получить текущую директорию: %v", err))
	}

	tempDir := filepath.Join(exeDir, "docker-build-temp")
	err = os.MkdirAll(tempDir, 0755)
	if err != nil {
		panic(fmt.Errorf("не удалось создать временную директорию: %v", err))
	}
	defer os.RemoveAll(tempDir)

	fileName := types.SelectFileExtension(object.Compiler)

	goFilePath := filepath.Join(tempDir, fileName)
	err = os.WriteFile(goFilePath, []byte(object.Code), 0644)
	if err != nil {
		panic(fmt.Errorf("не удалось создать Go файл: %v", err))
	}

	dockerfileContent := types.GetDockerfileContent(object.Compiler)

	dockerfilePath := filepath.Join(tempDir, "Dockerfile")
	err = os.WriteFile(dockerfilePath, []byte(dockerfileContent), 0644)
	if err != nil {
		panic(fmt.Errorf("не удалось создать Dockerfile: %v", err))
	}

	fmt.Printf("Временная директория: %s\n", tempDir)
	fmt.Printf("Go файл создан: %s\n", goFilePath)
	fmt.Printf("Dockerfile создан: %s\n", dockerfilePath)

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(fmt.Errorf("не удалось создать Docker клиент: %v", err))
	}

	buildContext, err := archive.TarWithOptions(tempDir, &archive.TarOptions{})
	if err != nil {
		panic(fmt.Errorf("не удалось создать архив для сборки: %v", err))
	}

	imageName := "go-task"
	buildResponse, err := cli.ImageBuild(ctx, buildContext, build.ImageBuildOptions{
		Tags:       []string{imageName},
		Remove:     true,
		Dockerfile: "Dockerfile",
	})
	if err != nil {
		panic(fmt.Errorf("ошибка сборки образа: %v", err))
	}
	defer buildResponse.Body.Close()

	_, err = io.Copy(os.Stdout, buildResponse.Body)
	if err != nil {
		panic(fmt.Errorf("ошибка чтения вывода сборки: %v", err))
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
	}, nil, nil, nil, "")
	if err != nil {
		panic(fmt.Errorf("ошибка создания контейнера: %v", err))
	}

	err = cli.ContainerStart(ctx, resp.ID, container.StartOptions{})
	if err != nil {
		panic(fmt.Errorf("ошибка запуска контейнера: %v", err))
	}

	var containerOutput bytes.Buffer

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(fmt.Errorf("ошибка ожидания контейнера: %v", err))
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: false,
		Timestamps: false,
		Follow:     false,
		Tail:       "all",
	})
	if err != nil {
		panic(fmt.Errorf("ошибка получения логов: %v", err))
	}
	defer out.Close()

	if _, err := stdcopy.StdCopy(&containerOutput, &containerOutput, out); err != nil {
		panic(fmt.Errorf("ошибка копирования логов: %v", err))
	}
	res := containerOutput.String()

	var containerErrorOutput bytes.Buffer

	errout, err := cli.ContainerLogs(ctx, resp.ID, container.LogsOptions{
		ShowStdout: false,
		ShowStderr: true,
		Timestamps: false,
		Follow:     false,
		Tail:       "all",
	})
	if err != nil {
		panic(fmt.Errorf("ошибка получения логов: %v", err))
	}
	defer errout.Close()

	if _, err := stdcopy.StdCopy(&containerErrorOutput, &containerErrorOutput, errout); err != nil {
		panic(fmt.Errorf("ошибка копирования логов: %v", err))
	}
	errres := containerErrorOutput.String()

	if err := cli.ContainerRemove(ctx, resp.ID, container.RemoveOptions{}); err != nil {
		fmt.Printf("Предупреждение: не удалось удалить контейнер: %v\n", err)
	}

	return &domain.ObjectResult{
		ID:       object.ID,
		Code:     object.Code,
		Compiler: object.Compiler,
		Result:   res,
		Error:    errres,
	}
}
