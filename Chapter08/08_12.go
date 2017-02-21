package main

import (
  "fmt"
  "time"
)

func main() {
  channel := make(chan string, 1)

  go func() {
    channel <- "Hello World!"
    println("Finishing goroutine")
  }()

  time.Sleep(time.Second)

  message := <-channel
  fmt.Println(message)
}
