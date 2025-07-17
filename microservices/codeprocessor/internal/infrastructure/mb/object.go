package mb

import domain "codeproc/microservices/codeprocessor/internal/domain/entity"

type TaskPublisher interface {
	Publish(object domain.ObjectResult) error
}

type TaskResultSubscriber interface {
	SubscribeResults()
}
