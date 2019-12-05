package server

import (
	"fmt"
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port int
}

func (c *Config) GetAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func NewConfig() (c *Config) {
	v := viper.New()
	v.AddConfigPath("configs")
	v.SetConfigName("api")

	if err := v.ReadInConfig(); err != nil {
		log.Fatal("couldn't load config:", err)
	}

	pflag.String("host", "localhost", "connection host")
	pflag.Int("port", 8081, "connection port")
	pflag.Parse()
	if err := v.BindPFlags(pflag.CommandLine); err != nil {
		log.Fatal("couldn't bind flags:", err)
	}

	if err := v.Unmarshal(&c); err != nil {
		log.Fatal("couldn't read config:", err)
	}

	return
}
