package main

import (
	"github.com/labstack/gommon/log"
	"github.com/vskut/go-temp/internal/server"
)

func main() {
	c := server.NewConfig(8080)

	if err := server.Start(c); err != nil {
		log.Fatal(err)
	}
}
