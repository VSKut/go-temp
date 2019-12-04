package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type Client struct {
	Config *Config
}

func NewClient(c *Config) *Client {
	s := &Client{
		Config: c,
	}

	return s
}

func (s *Client) Start() {

	raddr, err := net.ResolveTCPAddr(s.Config.Type, fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port))
	if err != nil {
		log.Fatal("ResolveTCPAddr error: ", err)
	}

	conn, err := net.DialTCP(s.Config.Type, nil, raddr)
	if err != nil {
		log.Fatal("DialTCP error: ", err)
	}

	defer func() {
		_ = conn.Close()
	}()

	log.Printf("Connected to %s:%d", s.Config.Host, s.Config.Port)

	fmt.Println("Type your message:")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">: ")
		rs, _ := reader.ReadString('\n')

		if rs == "\n" {
			continue
		}

		if rs == "exit\n" {
			os.Exit(0)
		}

		_, err := conn.Write([]byte(rs))
		if err != nil {
			log.Println("Message sending error: ", err)
			break
		}

		response := make([]byte, 1024)

		_, err = conn.Read(response)
		if err != nil {
			log.Println("Server response error: ", err.Error())
			os.Exit(1)
		}

		fmt.Println(string(response))
	}

}
