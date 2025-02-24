package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	HttpPort       int           `env:"HTTP_PORT, default=3000"`
	DefaultTimeout time.Duration `env:"DEFAULT_TIMEOUT, default=2s"`
}

func NewEnvConfig() *Config {
	ctx := context.Background()

	var c Config

	if err := envconfig.Process(ctx, &c); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing environment values: %v", err)
		panic("Error parsing environment values")
	}

	return &c
}
