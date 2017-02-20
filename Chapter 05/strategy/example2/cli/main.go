package main

import (
	"flag"
	"log"
	"os"

	"github.com/sayden/go-design-patterns/behavioral/strategy/example2/shapes"
)

var output = flag.String("output", "text", "The output to use between "+
	"'console' and 'image' file")

func main() {
	flag.Parse()

	activeStrategy, err := shapes.Factory(*output)
	if err != nil {
		log.Fatal(err)
	}

	switch *output {
	case shapes.TEXT_STRATEGY:
		activeStrategy.SetWriter(os.Stdout)
	case shapes.IMAGE_STRATEGY:
		w, err := os.Create("/tmp/image.jpg")
		if err != nil {
			log.Fatal("Error opening image")
		}
		defer w.Close()

		activeStrategy.SetWriter(w)
	}

	err = activeStrategy.Draw()
	if err != nil {
		log.Fatal(err)
	}
}
