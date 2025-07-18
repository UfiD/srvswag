package rabbitmq

import (
	"codeproc/microservices/codeprocessor/cmd/app/config"
	"fmt"

	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queueName  string
}

func NewRabbitMQPublisher(cfg config.RabbitMQPublisher) (*RabbitMQPublisher, error) {
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

	return &RabbitMQPublisher{
		connection: conn,
		channel:    ch,
		queueName:  cfg.QueueName,
	}, nil
}
