package queue

import (
	"fmt"
	"time"

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

var (
	queueTemplate = `amqp://%s:%s@%s:%s/`
)

func (q *Queue) setQueueURL() string {
	url := fmt.Sprintf(queueTemplate,
		q.internal.Cfg.Queue.User,
		q.internal.Cfg.Queue.Password,
		q.internal.Cfg.Queue.Host,
		q.internal.Cfg.Queue.Port,
	)

	q.internal.L.Info(url)

	return url
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
	var c rabbitmq.Consumer
	var err error
	for i := 1; i <= 5; i++ {
		c, err = rabbitmq.NewConsumer(q.setQueueURL(), rabbitmq.Config{})
		if err != nil {
			q.internal.L.Error(err.Error())
			time.Sleep(time.Second * time.Duration(i*2))
			continue
		}
		break
	}
	if err != nil {
		q.internal.L.Error(err.Error())
		return err
	}
	q.internal.L.Info("consumer set")

	q.consumer = &c

	return nil
}

// SetPublisher creates a new instance
// of RabbitMQ Publisher and insert into
// Queue struct
func (q *Queue) SetPublisher() error {
	var p *rabbitmq.Publisher
	var err error
	for i := 1; i <= 5; i++ {
		p, err = rabbitmq.NewPublisher(q.setQueueURL(), rabbitmq.Config{})
		if err != nil {
			q.internal.L.Error(err.Error())
			time.Sleep(time.Second * time.Duration(i*2))
			continue
		}
		break
	}

	if err != nil {
		return err
	}
	q.internal.L.Info("publisher set")
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
