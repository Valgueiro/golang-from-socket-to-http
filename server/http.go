package server

import (
	"fmt"
	"net"
	"strconv"
)

type Config struct {
	Port int
}

type HTTPServer struct {
	Config Config
}

func (s *HTTPServer) Start() error {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(s.Config.Port))
	if err != nil {
		return fmt.Errorf("could not Listen: %w", err)
	}

	fmt.Printf("Listening on port %d...\n", s.Config.Port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go s.handleConnection(conn)
	}
}

func (s *HTTPServer) handleConnection(conn net.Conn) {
	fmt.Println(conn)
}
