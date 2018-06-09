package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type ConfigTarget struct {
	Path string `json:"path"`
	Type string `json:"type"`
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

func main() {

	configFile := flag.String("config", "config.json", "Path to config file")
	flag.Parse()

	config, err := NewConfig(*configFile)

	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("%+v\n", config)
	}
}
