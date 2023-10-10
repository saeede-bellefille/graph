package server

import (
	"graph/socket"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server struct {
	address string
	echo    *echo.Echo
	client  *socket.Client
}

func New(address, upstream string) *Server {
	return &Server{
		address: address,
		echo:    echo.New(),
		client:  socket.NewClient(upstream),
	}
}

func (s *Server) Start() error {
	s.echo.POST("/send", s.send)
	return s.echo.Start(s.address)
}

func (s *Server) send(c echo.Context) error {
	err := s.client.Proxy(c.Request().Body)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return err
	}
	c.String(http.StatusOK, "OK\n")
	return nil
}
