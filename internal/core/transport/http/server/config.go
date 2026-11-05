package core_http_server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr            string        `envconfig:"ADDR" required:"true"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" required:"true"`
}

func newConfig() (Config, error) {
	var cfg Config

	if err := envconfig.Process("HTTP", &cfg); err != nil {
		return Config{}, fmt.Errorf("process server config: %w", err)
	}

	return cfg, nil
}

func NewConfigMust() Config {
	cfg, err := newConfig()
	if err != nil {
		err = fmt.Errorf("get HTTP server config: %w", err)
		panic(err)
	}

	return cfg
}
