package config

import (
	"sync"

	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog"
)

var (
	// config is the private instance loaded during API instantiation.
	config *Config

	// once is used to ensure the configuration is loaded only one time.
	once sync.Once
)

// Config is a set of configuration attributes that can be used for various
// component constructions within the API (i.e. datastore, server, etc.).
type Config struct {

	// Environment is the intended deployment environment for the API instance.
	Environment string `env:"ENVIRONMENT" envDefault:"local"`

	// Port is the listener port of the HTTP server.
	Port string `env:"PORT" envDefault:":8000"`

	// LogLevel is the level at which to log messages.
	LogLevel zerolog.Level `env:"LOG_LEVEL" envDefault:"1"`
}

// MustLoad will attempt to load the configuration from the environment and
// panic if unable to do so.
func MustLoad() *Config {
	once.Do(func() {
		config = mustLoad()
	})
	return config
}

// mustLoad is the underlying do-er of fetching the configuration attributes
// from the environment and deserializing them into a *Config instance.
func mustLoad() *Config {

	// cfg: initialize
	cfg := new(Config)

	// cfg: load from environment
	// note: as long as there aren't any non-defaulted, but required,
	// configuration attributes this library won't error.
	_ = env.Parse(cfg)

	// cfg: return
	return cfg

}
