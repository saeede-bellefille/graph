package processor

import (
	"fmt"
	"io"
)

type Processor struct {
	count uint64
	size  uint64
}

func New() *Processor {
	return &Processor{}
}

func (p *Processor) Process(conn io.ReadWriteCloser) error {
	defer conn.Close()
	buf, err := io.ReadAll(conn)
	if err != nil {
		return err
	}
	p.count++
	p.size += uint64(len(buf))
	fmt.Printf("%d: Received: %d, Total Size: %d\n", p.count, len(buf), p.size)
	return nil
}
