package main

import (
  "fmt"
  "strings"
)

func main() {
  toUpperSync("Hello Callbacks!", func(v string) {
    fmt.Printf("Callback: %s\n", v)
  })
}

func toUpperSync(word string, f func(string)) {
  f(strings.ToUpper(word))
}
