package handler

import (
	"bytes"
	"graph/socket"
	"io"
	"log"
)

type Handler struct {
	client *socket.Client
}

func New(upstream string) *Handler {
	return &Handler{
		client: socket.NewClient(upstream),
	}
}

func (h *Handler) Handle(conn io.ReadWriteCloser) error {
	defer conn.Close()
	buf, err := io.ReadAll(conn)
	if err != nil {
		return err
	}

	log.Print(string(buf))

	return h.client.Proxy(bytes.NewReader(buf))
}
