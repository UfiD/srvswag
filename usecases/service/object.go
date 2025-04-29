package service

import (
	"codeproc/infrastructure/consumer"
	"codeproc/infrastructure/repository"
	"codeproc/pkg/uuid"
)

type Object struct {
	repo     repository.Object
	consumer consumer.Object
}

func NewObject(repo repository.Object, consumer consumer.Object) *Object {
	return &Object{
		repo:     repo,
		consumer: consumer,
	}
}

func (uc *Object) Do(id string) {
	result := uc.consumer.Do()
	uc.repo.Put(id, result)
}

func (uc *Object) Post(code, compiler string) string {
	id := uuid.GetUUID()
	uc.repo.Post(id, code, compiler)
	go uc.Do(id)
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
