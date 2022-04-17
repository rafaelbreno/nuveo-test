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
		Queue    Queue
		Service  Service
		Database Database
	}
	// Queue stores values to
	// set Queue connection and
	// configs
	Queue struct {
		URL          string
		Name         string
		ConsumerName string
		User         string
		Password     string
		Port         string
		QueueName    string
	}

	// Database stores values to
	// set Database connections.
	// e.g Postgres, Redis, etc.
	Database struct {
		PGHost     string
		PGPort     string
		PGUser     string
		PGPassword string
		PGDBName   string
	}

	// Service stores value
	// of general use through
	// the services
	Service struct {
		NewClients string
		APIPort    string
		APIName    string
	}
)

// NewConfig set a Config instance
// based on environment variables
func NewConfig() (*Config, error) {
	if os.Getenv("ENV") != "docker" {
		err := godotenv.Load()
		if err != nil {
			return &Config{}, err
		}
	}

	return &Config{
		Queue: Queue{
			URL:          os.Getenv("RABBITMQ_URL"),
			QueueName:    os.Getenv("RABBITMQ_QUEUE_NAME"),
			ConsumerName: os.Getenv("RABBITMQ_CONSUMER_NAME"),
			User:         os.Getenv("RABBITMQ_USER"),
			Password:     os.Getenv("RABBITMQ_PASSWORD"),
			Port:         os.Getenv("RABBITMQ_PORT"),
			Name:         os.Getenv("RABBITMQ_NAME"),
		},
		Database: Database{
			PGHost:     os.Getenv("PGSQL_HOST"),
			PGPort:     os.Getenv("PGSQL_PORT"),
			PGUser:     os.Getenv("PGSQL_USER"),
			PGPassword: os.Getenv("PGSQL_PASSWORD"),
			PGDBName:   os.Getenv("PGSQL_DBNAME"),
		},
		Service: Service{
			NewClients: os.Getenv("NEW_CLIENTS"),
			APIPort:    os.Getenv("API_PORT"),
			APIName:    os.Getenv("API_NAME"),
		},
	}, nil
}
