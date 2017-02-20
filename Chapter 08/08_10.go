package main

import (
  "sync"
  "time"
)

type Counter struct {
  sync.Mutex
  value int
}

func main() {
  counter := Counter{}

  for i := 0; i < 10; i++ {
    go func(i int) {
      //race counter.Lock()
      counter.value++
      //counter.Unlock()
    }(i)
  }
  time.Sleep(time.Second)
  counter.Lock()
  println(counter.value)
  counter.Unlock()
}