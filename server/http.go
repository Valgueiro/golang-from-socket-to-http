package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

const CRLF = "\r\n"

type Config struct {
	Port         int
	EchoMessages bool
}

type HTTPServer struct {
	Config Config
}

func NewHttpServer(c Config) *HTTPServer {
	return &HTTPServer{
		Config: c,
	}
}

func (s *HTTPServer) Start() error {
	addr := &net.TCPAddr{
		Port: s.Config.Port,
	}

	ln, err := net.ListenTCP("tcp", addr)
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
	defer s.closeConnection(conn)
	reader := bufio.NewReader(conn)
	buff := ""
	msg := ""

	for buff != CRLF {
		fmt.Printf("waiting... Last buff: %q\n", buff)
		buff, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("client closed the connection")
				break
			} else {
				fmt.Println("read err: %w")
			}
		}
		fmt.Printf("Received: %q\n", buff)
		msg += buff

		if s.Config.EchoMessages {
			fmt.Printf("Echoing: %s\n", buff)

			conn.Write([]byte(buff + "\n"))
		}
	}

	fmt.Println("Message completed!")
}

func (s *HTTPServer) closeConnection(conn net.Conn) {
	fmt.Println("hello")
	conn.Write([]byte("Server closed the connection"))
	conn.Close()
}
