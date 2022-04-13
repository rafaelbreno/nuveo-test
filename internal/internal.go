package internal

import (
	"github.com/rafaelbreno/nuveo-test/config"
	"go.uber.org/zap"
)

type (
	// Internal stores variables/pointers
	// that will be used through the microservices.
	// E.g logging, config values, etc.
	Internal struct {
		Cfg *config.Config
		// L stands for Logger, just for
		// a cleaner code
		L *zap.Logger
	}
)

// NewInternal creates and instance for
// Internal struct.
func NewInternal(cfg *config.Config) (*Internal, error) {
	// There's no need to handle the returned error
	// because it's not being used a custom config
	l, _ := zap.NewProduction()
	return &Internal{
		L:   l,
		Cfg: cfg,
	}, nil
}
