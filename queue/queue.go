package queue

import (
	rabbitmq "github.com/wagslane/go-rabbitmq"
)

type (
	// Queue abstracts the RabbitMQ library
	// for easier use through the app.
	Queue struct {
		consumer  rabbitmq.Consumer
		publisher rabbitmq.Publisher
	}
)

// NewQueue set a Queue instance
// given a Config
func NewQueue() (*Queue, error) {
	return &Queue{}, nil
}
