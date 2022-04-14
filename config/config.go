package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	// Config stores values that
	// is used through the service.
	// E.g set RabbitMQ connection
	// E.g set API port
	Config struct {
		Queue   Queue
		Service Service
	}
	// Queue stores values to
	// set Queue connection and
	// configs
	Queue struct {
		URL          string
		Name         string
		ConsumerName string
	}

	// Service stores value
	// of general use through
	// the services
	Service struct {
		NewClients string
	}
)

// NewConfig set a Config instance
// based on environment variables
func NewConfig() (*Config, error) {
	if os.Getenv("ENV") != "" {
		err := godotenv.Load()
		if err != nil {
			return &Config{}, err
		}
	}

	return &Config{
		Queue: Queue{
			URL:          os.Getenv("RABBITMQ_URL"),
			Name:         os.Getenv("RABBITMQ_QUEUE_NAME"),
			ConsumerName: os.Getenv("RABBITMQ_CONSUMER_NAME"),
		},
		Service: Service{
			NewClients: os.Getenv("NEW_CLIENTS"),
		},
	}, nil
}
