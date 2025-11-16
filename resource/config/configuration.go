package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Service struct {
		Port struct {
			REST string `yaml:"rest"`
		} `yaml:"port"`
	} `yaml:"service"`
}

func NewConfiguration(path string) (*Configuration, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal into struct
	var cfg Configuration
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
