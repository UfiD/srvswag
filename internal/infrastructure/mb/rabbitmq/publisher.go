package rabbitmq

import (
	"codeproc/cmd/app/config"
	"codeproc/internal/domain"
	"encoding/json"
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

func (rbp *RabbitMQPublisher) Publish(object domain.Object) error {
	body, err := json.Marshal(object)
	if err != nil {
		return err
	}

	err = rbp.channel.Publish(
		"",
		rbp.queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}
	return nil
}

func (rbp *RabbitMQPublisher) Close() {
	rbp.channel.Close()
	rbp.connection.Close()
}
