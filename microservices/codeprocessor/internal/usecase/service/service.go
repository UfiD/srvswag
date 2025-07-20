package service

import (
	"codeproc/microservices/codeprocessor/internal/domain"
	"codeproc/microservices/codeprocessor/internal/infrastructure/mb"
)

type usecase struct {
	rmp mb.TaskResultPublisher
	rms mb.TaskSubscriber
	d   domain.Object
}

func NewUsecase(rmp mb.TaskResultPublisher,
	rms mb.TaskSubscriber,
	d domain.Object) *usecase {
	return &usecase{
		rmp: rmp,
		rms: rms,
		d:   d,
	}
}

func (uc *usecase) Start() {
	uc.rms.SubscribeResults()
}
