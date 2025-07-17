package main

import (
	"codeproc/cmd/app/config"
	_ "codeproc/docs"
	controller "codeproc/internal/controller/http"
	"codeproc/internal/infrastructure/mb/rabbitmq"
	"codeproc/internal/infrastructure/repository/ram_storage"
	sessionstorage "codeproc/internal/infrastructure/repository/session_storage"
	"codeproc/internal/usecases/service"
	"codeproc/internal/usecases/session"
	pkgHttp "codeproc/pkg/http"
	"log"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title CodeProcessor
// @version 1.0
// @description This is a HTTP-server for code processing.

// @host localhost:8080
// @BasePath /task

func main() {
	appFlags := config.ParseFlag()
	var cfg config.AppConfig
	config.ParseConfig(appFlags.ConfigPath, &cfg)
	storage := ram_storage.NewStorage()
	rabbitMQPublisher, err := rabbitmq.NewRabbitMQPublisher(cfg.RabbitMQPublisher)
	if err != nil {
		panic(err)
	}
	service := service.NewObject(storage, rabbitMQPublisher)
	userstore := sessionstorage.NewObject()
	sessionstore := sessionstorage.NewSessionStorage()
	manager := session.NewObject(userstore, sessionstore, 3600)
	server := controller.New(service, manager)

	r := chi.NewRouter()
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	server.WithObjectHandler(r)

	if err := pkgHttp.CreateAndRunServer(cfg.HTTPConfig.Addr, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
