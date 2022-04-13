package config

type (
	// Config stores values that
	// is used through the service.
	// E.g set RabbitMQ connection
	// E.g set API port
	Config struct {
		Queue Queue
	}
	// Queue stores values to
	// set Queue connection and
	// configs
	Queue struct {
		URL  string
		Name string
	}
)

// NewConfig set a Config instance
// based on environment variables
func NewConfig() (*Config, error) {
	return &Config{}, nil
}
