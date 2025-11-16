package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Credential struct {
	Database struct {
		Host     string `yaml:"host"`
		Name     string `yaml:"name"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		Driver   string `yaml:"driver"`
	}
}

func NewCredential(path string) (*Credential, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal into struct
	var credential Credential
	if err = yaml.Unmarshal(data, &credential); err != nil {
		return nil, err
	}

	return &credential, nil
}
