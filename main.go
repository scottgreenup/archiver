package main

import (
	"flag"
	"fmt"
	"os"
)

type Archiver struct {
	FromPath string
	ToPath string
}

func NewArchiver(target ConfigTarget) (*Archiver, error) {
	from, err := AbsPath(target.From)
	if err != nil {
		return nil, err
	}
	if err := ValidateFrom(from); err != nil {
		return nil, err
	}

	to, err := AbsPath(target.To)
	if err != nil {
		return nil, err
	}
	if err := ValidateTo(to); err != nil {
		return nil, err
	}

	return &Archiver{
		FromPath: from,
		ToPath: to,
	}, nil

}

func main() {
	configFile := flag.String("config", "Archiverfile", "Path to config file")
	flag.Parse()

	config, err := NewConfig(*configFile)

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("%+v\n", config)
	}

	a, err := NewArchiver(config.Targets["GPG"])
	fmt.Printf("%+v %+v\n", a, err)
}
