package main

import (
	"github.com/rafaelbreno/nuveo-test/config"
	"github.com/rafaelbreno/nuveo-test/internal"
	"github.com/rafaelbreno/nuveo-test/queue"
	"github.com/rafaelbreno/nuveo-test/services/api/server"
	"github.com/rafaelbreno/nuveo-test/services/api/storage"
	"go.uber.org/zap"
)

func main() {
	l, _ := zap.NewProduction()

	cfg, err := config.NewConfig()
	if err != nil {
		l.Fatal(err.Error())
	}

	in := internal.NewInternal(cfg)

	q := queue.NewQueue(in)

	q.SetPublisher()

	st, err := storage.NewSQL(in)
	if err != nil {
		l.Fatal(err.Error())
	}

	server.NewServer(in, &storage.Storage{
		SQL: st,
	}, q).Start()
}
