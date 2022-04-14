package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rafaelbreno/nuveo-test/config"
	"github.com/rafaelbreno/nuveo-test/internal"
	"github.com/rafaelbreno/nuveo-test/queue"
	"github.com/wagslane/go-rabbitmq"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	in := internal.NewInternal(cfg)

	q := queue.NewQueue(in)

	if err := q.SetConsumer(); err != nil {
		in.L.Fatal(err.Error())
	}

	defer func() {
		q.Consumer().Close()
	}()

	if err := q.
		Consumer().
		StartConsuming(func(d rabbitmq.Delivery) rabbitmq.Action {
			// implement logic
			return rabbitmq.Ack
		},
			in.Cfg.Queue.Name,
			[]string{},
			rabbitmq.WithConsumeOptionsConcurrency(10),
			rabbitmq.WithConsumeOptionsQueueDurable,
			rabbitmq.WithConsumeOptionsQuorum,
			rabbitmq.WithConsumeOptionsBindingExchangeName("events"),
			rabbitmq.WithConsumeOptionsBindingExchangeKind("topic"),
			rabbitmq.WithConsumeOptionsBindingExchangeDurable,
			rabbitmq.WithConsumeOptionsConsumerName(in.Cfg.Queue.ConsumerName),
		); err != nil {
		in.L.Fatal(err.Error())
	}

	// block main thread - wait for shutdown signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		in.L.Info(sig.String())
		done <- true
	}()

	in.L.Info("awaiting signal")
	<-done
	in.L.Info("stopping consumer")
}
