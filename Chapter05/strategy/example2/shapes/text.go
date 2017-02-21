package shapes

import "github.com/sayden/go-design-patterns/behavioral/strategy/example2"

type TextSquare struct {
	strategy.DrawOutput
}

func (t *TextSquare) Draw() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}
