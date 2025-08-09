package main

import (
	"github.com/Valgueiro/golang-from-socket-to-http/server"
)

func main() {
	s := &server.HTTPServer{
		Config: server.Config{
			Port: 8080,
		},
	}

	s.Start()
}
