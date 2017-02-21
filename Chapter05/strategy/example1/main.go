package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

type OutputStrategy interface {
	Draw() error
}

type TextSquare struct{}

func (t *TextSquare) Draw() error {
	println("Circle")
	return nil
}

type ImageSquare struct {
	DestinationFilePath string
}

func (t *ImageSquare) Draw() error {
	width := 800
	height := 600

	bgColor := image.Uniform{color.RGBA{R: 70, G: 70, B: 70, A:0}}
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

	w, err := os.Create(t.DestinationFilePath)
	if err != nil {
		return fmt.Errorf("Error opening image")
	}
	defer w.Close()

	if err = jpeg.Encode(w, bgRectangle, quality); err != nil {
		return fmt.Errorf("Error writing image to disk")
	}

	return nil
}

var output = flag.String("output", "console", "The output to use between 'console' and 'image' file")

func main() {
	flag.Parse()

	var activeStrategy OutputStrategy

	switch *output {
	case "console":
		activeStrategy = &TextSquare{}
	case "image":
		activeStrategy = &ImageSquare{"/tmp/image.jpg"}
	default:
		activeStrategy = &TextSquare{}
	}

	err := activeStrategy.Draw()
	if err != nil {
		log.Fatal(err)
	}
}
