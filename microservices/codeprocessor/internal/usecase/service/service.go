package service

import (
	"codeproc/microservices/codeprocessor/internal/domain"
	"codeproc/microservices/codeprocessor/internal/infrastructure/mb"
)

type usecase struct {
	rmp mb.TaskPublisher
	rms mb.TaskResultSubscriber
	d   domain.Object
}

func NewUsecase(rmp mb.TaskPublisher,
	rms mb.TaskResultSubscriber,
	d domain.Object) *usecase {
	return &usecase{
		rmp: rmp,
		rms: rms,
		d:   d,
	}
}

func (uc *usecase) Start() {
	for {

	}
}

func (uc *usecase) Consume() {

}
