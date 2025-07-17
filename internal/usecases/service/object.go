package service

import (
	"codeproc/internal/domain"
	"codeproc/internal/infrastructure/mb"
	"codeproc/internal/infrastructure/repository"
	"codeproc/pkg/uuid"
)

type Object struct {
	repo      repository.Object
	publisher mb.TaskPublisher
}

func NewObject(repo repository.Object, publisher mb.TaskPublisher) *Object {
	return &Object{
		repo:      repo,
		publisher: publisher,
	}
}

func (uc *Object) Post(object domain.Object) string {
	id := uuid.GetUUID()
	uc.repo.Create(id, object.Code, object.Compiler)
	object.ID = id
	uc.publisher.Publish(object)
	return id
}

func (uc *Object) GetStatus(id string) (string, error) {
	status, err := uc.repo.GetStatus(id)
	return status, err
}

func (uc *Object) GetResult(id string) (string, error) {
	res, err := uc.repo.GetResult(id)
	return res, err
}
