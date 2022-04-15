package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/rafaelbreno/nuveo-test/config"
	"github.com/rafaelbreno/nuveo-test/internal"
	"github.com/rafaelbreno/nuveo-test/services/writer/writer"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	in := internal.NewInternal(cfg)

	w := writer.NewWriter(in)

	defer func() {
		w.CloseConsumer()
	}()

	if err := w.StartConsuming(); err != nil {
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
