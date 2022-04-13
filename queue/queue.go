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
// given a Config
func NewQueue(in *internal.Internal) (*Queue, error) {
	c, err := rabbitmq.NewConsumer(in.Cfg.Queue.URL, rabbitmq.Config{})
	if err != nil {
		in.L.Error(err.Error())
		return &Queue{}, err
	}

	p, err := rabbitmq.NewPublisher(in.Cfg.Queue.URL, rabbitmq.Config{})
	if err != nil {
		in.L.Error(err.Error())
		return &Queue{}, err
	}

	return &Queue{
		consumer:  &c,
		publisher: p,
		internal:  in,
	}, nil
}
