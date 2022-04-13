package internal

import (
	"testing"

	"github.com/rafaelbreno/nuveo-test/config"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestInternal(t *testing.T) {
	assert := assert.New(t)

	{
		cfg := &config.Config{
			Queue: config.Queue{
				URL:  "url",
				Name: "queue",
			},
		}

		l, _ := zap.NewProduction()

		wantInternal := &Internal{
			Cfg: cfg,
			L:   l,
		}
		gotInternal := NewInternal(cfg)

		// Not possible to assert Logger
		// because of atomic level.
		assert.Equal(*wantInternal.Cfg, *gotInternal.Cfg)
	}

}
