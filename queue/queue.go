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

// Consumer returns the rabbitmq.Consumer
// pointer.
func (q *Queue) Consumer() *rabbitmq.Consumer {
	return q.consumer
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

// Publisher returns the rabbitmq.Publisher
// pointer.
func (q *Queue) Publisher() *rabbitmq.Publisher {
	return q.publisher
}

// Publish sends data to Queue.
func (q *Queue) Publish(data []byte, key ...string) error {
	return q.publisher.Publish(
		data,
		key,
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsMandatory,
		rabbitmq.WithPublishOptionsPersistentDelivery,
		rabbitmq.WithPublishOptionsExchange("events"),
	)
}

// PublishCreate is a shortcut
// for Publish(data, "create")
func (q *Queue) PublishCreate(data []byte) error {
	return q.Publish(data, "create")
}

// PublishDelete is a shortcut
// for Publish(data, "delete")
func (q *Queue) PublishDelete(data []byte) error {
	return q.Publish(data, "delete")
}
