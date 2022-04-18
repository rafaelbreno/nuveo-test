package server

import (
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/nuveo-test/internal"
	"github.com/rafaelbreno/nuveo-test/queue"
	"github.com/rafaelbreno/nuveo-test/services/api/handler"
	"github.com/rafaelbreno/nuveo-test/services/api/server/middleware"
	"github.com/rafaelbreno/nuveo-test/services/api/storage"
)

type (
	// Server handles all HTTP requests
	Server struct {
		HTTP  *fiber.App
		in    *internal.Internal
		st    *storage.Storage
		queue *queue.Queue
	}
)

// NewServer creates Server instance
// based on given internal.
func NewServer(in *internal.Internal, st *storage.Storage, queue *queue.Queue) *Server {
	return &Server{
		HTTP: fiber.New(fiber.Config{
			CaseSensitive: true,
			StrictRouting: false,
			Concurrency:   256 * 1024,
			WriteTimeout:  time.Duration(45 * time.Second),
		}),
		in:    in,
		st:    st,
		queue: queue,
	}
}

func (s *Server) Start() {
	r := s.HTTP.Use(middleware.CheckBody)

	uh := handler.NewUserHandler(s.st, s.in, s.queue)

	r.Get("user/:uuid", uh.Read)
	r.Get("users/", uh.ReadAll)
	r.Post("user/create", uh.Create)
	r.Put("user/:uuid", uh.Update)
	r.Patch("user/:uuid", uh.Update)
	r.Delete("user/:uuid", uh.Delete)

	s.HTTP.Listen(":" + s.in.Cfg.Service.APIPort)
}
