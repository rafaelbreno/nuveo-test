package server

import (
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/rafaelbreno/nuveo-test/internal"
)

type (
	// Server handles all HTTP requests
	Server struct {
		HTTP *fiber.App
		in   *internal.Internal
	}
)

// NewServer creates Server instance
// based on given internal.
func NewServer(in *internal.Internal) *Server {
	return &Server{
		HTTP: fiber.New(fiber.Config{
			CaseSensitive: true,
			StrictRouting: false,
			Concurrency:   256 * 1024,
			WriteTimeout:  time.Duration(45 * time.Second),
		}),
		in: in,
	}
}
