package writer

import (
	"encoding/json"
	"fmt"

	"github.com/rafaelbreno/nuveo-test/entity"
	"github.com/rafaelbreno/nuveo-test/internal"
	"github.com/rafaelbreno/nuveo-test/queue"
	"github.com/rafaelbreno/nuveo-test/services/writer/storage"
	"github.com/wagslane/go-rabbitmq"
)

type (
	// Writer stores all values
	// necessery to run the microservice.
	Writer struct {
		In      *internal.Internal
		Queue   *queue.Queue
		Storage storage.Storage
	}
)

// NewWriter builds and instance of Writer.
func NewWriter(in *internal.Internal) Writer {
	q := queue.NewQueue(in)

	if err := q.SetConsumer(); err != nil {
		in.L.Error(err.Error())
	}

	f, err := storage.NewFile(in)
	if err != nil {
		in.L.Fatal(err.Error())
	}

	return Writer{
		In:      in,
		Queue:   q,
		Storage: &f,
	}
}

// CloseConsumer gracefully disconnects
// from RabbitMQ.
func (w *Writer) CloseConsumer() error {
	return w.Queue.Consumer().Close()
}

// StartConsuming start receiving RabbitMQ Deliveries
// and processing them.
func (w *Writer) StartConsuming() error {
	return w.
		Queue.
		Consumer().
		StartConsuming(
			w.writeEntry,
			w.In.Cfg.Queue.Name,
			[]string{},
			rabbitmq.WithConsumeOptionsConcurrency(10),
			rabbitmq.WithConsumeOptionsQueueDurable,
			rabbitmq.WithConsumeOptionsQuorum,
			rabbitmq.WithConsumeOptionsBindingExchangeName("events"),
			rabbitmq.WithConsumeOptionsBindingExchangeKind("topic"),
			rabbitmq.WithConsumeOptionsBindingExchangeDurable,
			rabbitmq.WithConsumeOptionsConsumerName(w.In.Cfg.Queue.ConsumerName),
		)
}

// writeEntry receives rabbitmq.Delivery instance
// and starts processing it.
// Depending on the results, the file can be stored
// or sent to a dead-letter queue.
func (w *Writer) writeEntry(d rabbitmq.Delivery) rabbitmq.Action {
	w.In.L.Info(fmt.Sprintf("processing message %s", d.MessageId))

	u := entity.User{}

	if err := json.Unmarshal(d.Body, &u); err != nil {
		w.In.L.Error(err.Error())
		// sending to a dead-letter queue
		return rabbitmq.NackDiscard
	}

	if err := w.Storage.MarshalAndWrite(u, u.UUID); err != nil {
		w.In.L.Error(err.Error())
		// sending to a dead-letter queue
		return rabbitmq.NackDiscard
	}

	return rabbitmq.Ack
}
