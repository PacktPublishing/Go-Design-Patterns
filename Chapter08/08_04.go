package main
import "time"

func main() {
  messagePrinter := func(msg string) {
    println(msg)
  }

  go messagePrinter("Hello World")
  go messagePrinter("Hello goroutine")
  time.Sleep(time.Second)
}

