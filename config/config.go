package config

import (
	"time"

	"github.com/joeshaw/envdecode"
)

// Conf is an App configuration struct
type Conf struct {
	Debug  bool `env:"DEBUG,required"`
	Server serverConf
}

type serverConf struct {
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
}

// AppConfig populates a Conf app configuration object
func AppConfig() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		//	main.Logger.Fatal().Err(err).Msg("Failed to decode: %s", err)
	}
	return &c
}
