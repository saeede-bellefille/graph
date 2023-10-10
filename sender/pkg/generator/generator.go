package generator

import (
	"bytes"
	"math/rand"
	"net/http"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Generator struct {
	minLength int
	maxLength int
	upstream  string
}

func New(minLength, maxLength int, upstream string) *Generator {
	return &Generator{
		minLength: minLength,
		maxLength: maxLength,
		upstream:  upstream,
	}
}

func (g *Generator) Send() error {
	count := rand.Intn(g.maxLength-g.minLength) + g.minLength
	data := randData(count)
	resp, err := http.Post(g.upstream+"/send", "text/plain", bytes.NewReader(data))
	if err != nil {
		return err
	}
	resp.Body.Close()
	return nil
}

func randData(count int) []byte {
	b := make([]byte, count)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return b
}
