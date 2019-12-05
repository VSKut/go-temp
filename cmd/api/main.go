package main

import (
	"github.com/labstack/gommon/log"
	"github.com/vskut/go-temp/internal/server"
)

func main() {
	if err := server.Start(server.NewConfig()); err != nil {
		log.Fatal(err)
	}
}
