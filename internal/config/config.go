package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port              string `envconfig:"PORT" default:"8000"`
	DSN               string `envconfig:"DSN" required:"true"`
	ReadTimeout       int    `envconfig:"READ_TIMEOUT" default:"30" required:"true"`
	WriteTimeout      int    `envconfig:"WRITE_TIMEOUT" default:"30" required:"true"`
	ReadHeaderTimeout int    `envconfig:"READ_HEADER_TIMEOUT" default:"30" required:"true"`
}

func GetConfig() (Config, error) {
	var isLocal bool
	flag.BoolVar(&isLocal, "local", false, "Use env vars from 'local.env' (should be in root)")
	flag.Parse()
	if isLocal {
		log.Println("using local config")
		if err := godotenv.Load("local.env"); err != nil {
			return Config{}, fmt.Errorf("error loading 'local.env' file: %w", err)
		}
	}
	config := Config{}

	err := envconfig.Process("", &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
