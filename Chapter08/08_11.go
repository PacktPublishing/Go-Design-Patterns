package main

import (
  "fmt"
  "time"
)

func main() {
  channel := make(chan string)

  var waitGroup sync.WaitGroup

  waitGroup.Add(1)
  go func() {
    channel <- "Hello World!"
    println("Finishing goroutine")
    waitGroup.Done()
  }()

  time.Sleep(time.Second)
  message := <-channel
  fmt.Println(message)
  waitGroup.Wait()
}

