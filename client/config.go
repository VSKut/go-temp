package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Host string
	Port int
	Type string
}

func NewConfig() (c *Config) {
	v := viper.New()
	v.AddConfigPath("server")
	v.SetConfigName("config")

	if err := v.ReadInConfig(); err != nil {
		log.Fatal("couldn't load config:", err)
	}

	pflag.String("host", "localhost", "connection host")
	pflag.Int("port", 8081, "connection port")
	pflag.String("type", "tcp", "connection type")
	pflag.Parse()
	if err := v.BindPFlags(pflag.CommandLine); err != nil {
		log.Fatal("couldn't bind flags:", err)
	}

	if err := v.Unmarshal(&c); err != nil {
		log.Fatal("couldn't read config:", err)
	}

	return
}
