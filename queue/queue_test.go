package queue

import (
	"testing"

	"github.com/rafaelbreno/nuveo-test/config"
	"github.com/rafaelbreno/nuveo-test/internal"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

const (
	rabbitURL   string = "amqp://localhost:5672/%2f"
	rabbitQueue string = "queue"
)

func TestQueue(t *testing.T) {
	assert := assert.New(t)

	cfg := zap.NewDevelopmentConfig()
	cfg.Level.SetLevel(zap.FatalLevel)
	l, _ := cfg.Build()

	{
		want := &config.Config{
			Queue: config.Queue{
				URL:  "",
				Name: "",
			},
		}

		got := NewQueue(&internal.Internal{
			Cfg: want,
		})
		assert.Equal(want, got.internal.Cfg)
	}

	{
		in := internal.Internal{
			Cfg: &config.Config{
				Queue: config.Queue{
					URL:  "",
					Name: "",
				},
			},
			L: l,
		}

		q := NewQueue(&in)

		got := q.SetConsumer()

		assert.NotNil(got)
	}

	{
		in := internal.Internal{
			Cfg: &config.Config{
				Queue: config.Queue{
					URL:  "",
					Name: "",
				},
			},
			L: l,
		}

		q := NewQueue(&in)

		got := q.SetPublisher()

		assert.NotNil(got)
	}

}
