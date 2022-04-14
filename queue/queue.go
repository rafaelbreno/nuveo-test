package queue

import (
	"github.com/rafaelbreno/nuveo-test/internal"
	rabbitmq "github.com/wagslane/go-rabbitmq"
)

type (
	// Queue abstracts the RabbitMQ library
	// for easier use through the app.
	Queue struct {
		consumer  *rabbitmq.Consumer
		publisher *rabbitmq.Publisher
		internal  *internal.Internal
	}
)

// NewQueue set a Queue instance
// given an Internal struct
func NewQueue(in *internal.Internal) *Queue {
	return &Queue{
		internal: in,
	}
}

// SetConsumer creates a new instance
// of RabbitMQ Consumer and insert into
// Queue struct
func (q *Queue) SetConsumer() error {
	c, err := rabbitmq.NewConsumer(q.internal.Cfg.Queue.URL, rabbitmq.Config{})
	if err != nil {
		q.internal.L.Error(err.Error())
		return err
	}
	q.consumer = &c
	return nil
}

// SetPublisher creates a new instance
// of RabbitMQ Publisher and insert into
// Queue struct
func (q *Queue) SetPublisher() error {
	p, err := rabbitmq.NewPublisher(q.internal.Cfg.Queue.URL, rabbitmq.Config{})
	if err != nil {
		q.internal.L.Error(err.Error())
		return err
	}
	q.publisher = p
	return nil
}
