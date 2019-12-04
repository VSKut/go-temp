package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	Config *Config
}

func NewServer(c *Config) *Server {
	s := &Server{
		Config: c,
	}

	return s
}

func (s *Server) Start() {
	l, err := net.Listen(s.Config.Type, fmt.Sprintf("%s:%d", s.Config.Host, s.Config.Port))
	if err != nil {
		log.Fatal("TCP Server listener error:", err)
	}

	defer func() {
		_ = l.Close()
	}()

	log.Printf("Listening on %s:%d", s.Config.Host, s.Config.Port)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go handleMessage(conn)
	}

}

func handleMessage(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for {
		ok := scanner.Scan()

		if !ok {
			break
		}

		responseMessage := parseMessage(scanner.Text())
		_, err := conn.Write([]byte(responseMessage))
		if err != nil {
			log.Fatal("Response error:", err)
		}
	}
}

func parseMessage(m string) string {
	r, err := strconv.ParseInt(m, 10, 64)
	if err == nil {
		return strconv.Itoa(int(r * 2))
	}

	return strings.ToUpper(m)
}
