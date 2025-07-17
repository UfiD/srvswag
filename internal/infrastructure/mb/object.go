package mb

import (
	"codeproc/internal/domain"
)

type TaskPublisher interface {
	Publish(object domain.Object) error
}

type TaskResultSubscriber interface {
	SubscribeResults()
}
