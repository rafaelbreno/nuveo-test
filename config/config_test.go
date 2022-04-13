package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	envContent string = `
ENV=local

# RabbitMQ
RABBITMQ_URL=amqp://guest:guest@localhost:5672/
RABBITMQ_QUEUE_NAME=queue
	`
	rabbitURL   string = "amqp://guest:guest@localhost:5672/"
	rabbitQueue string = "queue"
)

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	{
		gotCfg, gotErr := NewConfig()
		wantCfg := &Config{
			Queue: Queue{
				URL:  "",
				Name: "",
			},
		}

		assert.Equal(wantCfg, gotCfg)
		assert.Nil(gotErr)
	}

	{
		os.Setenv("ENV", "local")
		_, gotErr := NewConfig()

		assert.NotNil(gotErr)

	}

	f, err := writeMockEnv()

	if err != nil {
		t.Fatal(err.Error())
	}

	defer f.Close()
	defer os.Remove(".env")

	{
		os.Setenv("ENV", "local")
		gotCfg, gotErr := NewConfig()
		wantCfg := &Config{
			Queue: Queue{
				URL:  rabbitURL,
				Name: rabbitQueue,
			},
		}

		assert.Equal(wantCfg, gotCfg)
		assert.Nil(gotErr)
		os.Setenv("ENV", "")
	}

}

func writeMockEnv() (*os.File, error) {
	f, err := os.Create(".env")
	if err != nil {
		return nil, err
	}

	f.WriteString(envContent)
	return f, nil
}
