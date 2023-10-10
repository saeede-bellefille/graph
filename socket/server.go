package socket

import (
	"io"
	"net"
)

type Server struct {
	address string
	handler func(io.ReadWriteCloser) error
}

func NewServer(address string, handler func(io.ReadWriteCloser) error) *Server {
	return &Server{
		address: address,
		handler: handler,
	}
}

func (s *Server) Start() error {
	listen, err := net.Listen("tcp", s.address)
	if err != nil {
		return err
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}
		go s.handler(conn)
	}
}
