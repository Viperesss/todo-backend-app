package core_logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	level  string `envconfig:"LEVEL"  required:"true"`
	folder string `envconfig:"FOLDER" required:"true"`
}

func newConfig() (Config, error) {
	var config Config

	// читает переменные окружения с префиксом LOGGER
	if err := envconfig.Process("LOGGER", &config); err != nil {
		return Config{}, fmt.Errorf("process env config: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := newConfig()
	if err != nil {
		err = fmt.Errorf("get Logger config: %w", err)
		panic(err)
	}

	return config
}
