package pubsub

import (
	"fmt"
	"io"
	"os"
	"time"
)

type writerSubscriber struct {
	in     chan interface{}
	id     int
	Writer io.Writer
}

func NewWriterSubscriber(id int, out io.Writer) Subscriber {
	if out == nil {
		out = os.Stdout
	}
	s := &writerSubscriber{
		id:     id,
		in:     make(chan interface{}),
		Writer: out,
	}
	go func() {
		for msg := range s.in {
			fmt.Fprintf(s.Writer, "(W%d): %v\n", s.id, msg)
		}
	}()
	return s
}

func (s *writerSubscriber) Notify(msg interface{}) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%#v", rec)
		}
	}()
	select {
	case s.in <- msg:
	case <-time.After(time.Second):
		err = fmt.Errorf("Timeout\n")
	}
	return
}

func (s *writerSubscriber) Close() {
	close(s.in)
}
