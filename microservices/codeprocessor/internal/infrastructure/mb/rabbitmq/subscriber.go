package rabbitmq

import (
	"codeproc/microservices/codeprocessor/cmd/app/config"
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQSubscriber struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queueName  string
}

func NewRabbitMQSubscriber(cfg config.RabbitMQSubscriber) (*RabbitMQSubscriber, error) {
	url := fmt.Sprintf("amqp://guest:guest@%s:%d", cfg.Host, cfg.Port)
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("connecting to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(
		cfg.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQSubscriber{
		connection: conn,
		channel:    ch,
		queueName:  cfg.QueueName,
	}, nil
}
