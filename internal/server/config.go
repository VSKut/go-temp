package server

import "fmt"

type Config struct {
	Port string
}

func NewConfig(port int) *Config {
	return &Config{
		Port: fmt.Sprintf(":%d", port),
	}
}
