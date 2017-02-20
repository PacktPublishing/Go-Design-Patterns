package main
import (
  "fmt"
  "strings"
  "sync"
)

var wait sync.WaitGroup

func main() {
  wait.Add(1)

  toUpperAsync("Hello Callbacks!", func(v string) {
    fmt.Printf("Callback: %s\n", v)
    wait.Done()
  })

  println("Waiting async response...")
  wait.Wait()
}

func toUpperAsync(word string, f func(string)) {
  go func(){
    f(strings.ToUpper(word))
  }()
}
