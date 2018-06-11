package main

import (
	"github.com/pkg/errors"

	"encoding/json"
	"os"
)

type ConfigTarget struct {
	From string     `json:"from"`
	To string       `json:"to"`
}

type Config struct {
	Targets map[string]ConfigTarget `json:"targets"`
}

func NewConfig(filename string) (*Config, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer jsonFile.Close()

	config := new(Config)
	if err := json.NewDecoder(jsonFile).Decode(config); err != nil {
		return nil, errors.WithStack(err)
	}
	return config, nil
}
