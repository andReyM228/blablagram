package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type (
	Config struct {
		Mongo MongoDB `yaml:"mongodb"`
		Hash  Hash    `yaml:"hash"`
	}

	MongoDB struct {
		Url string `yaml:"url"`
	}

	Hash struct {
		Salt string `yaml:"salt"`
	}
)

func ParseConfig() (Config, error) {
	file, err := os.ReadFile("./cmd/config.yaml")
	if err != nil {
		log.Fatalf("reading config: %s", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("unmarshal config: %s", err)
	}

	return cfg, nil
}
