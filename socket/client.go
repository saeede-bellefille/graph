package socket

import (
	"io"
	"net"
)

type Client struct {
	address string
}

func NewClient(address string) *Client {
	return &Client{
		address: address,
	}
}

func (c *Client) Proxy(input io.Reader) error {
	// TODO: use connection pool instead
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = io.Copy(conn, input)
	return err
}
