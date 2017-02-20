package shapes

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"

	"github.com/sayden/go-design-patterns/behavioral/strategy/example2"
)

type ImageSquare struct {
	strategy.DrawOutput
}

func (i *ImageSquare) Draw() error {
	width := 800
	height := 600

	bgColor := image.Uniform{color.RGBA{0, 0, 0, 1}}
	origin := image.Point{0, 0}
	quality := &jpeg.Options{Quality: 75}

	bgRectangle := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})

	draw.Draw(bgRectangle, bgRectangle.Bounds(), &bgColor, origin, draw.Src)

	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareImg := image.NewRGBA(square)

	draw.Draw(bgRectangle, squareImg.Bounds(), &squareColor, origin, draw.Src)

	if i.Writer == nil {
		return fmt.Errorf("No writer stored on ImageSquare")
	}
	if err := jpeg.Encode(i.Writer, bgRectangle, quality); err != nil {
		return fmt.Errorf("Error writing image to disk")
	}

	if i.LogWriter != nil {
		i.LogWriter.Write([]byte("Image written in provided writer\n"))
	}

	return nil
}
