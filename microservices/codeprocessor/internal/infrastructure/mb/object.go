package mb

import domain "codeproc/microservices/codeprocessor/internal/domain/entity"

type TaskResultPublisher interface {
	Publish(object domain.ObjectResult) error
}

type TaskSubscriber interface {
	SubscribeResults()
}
