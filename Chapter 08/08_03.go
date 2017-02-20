package main
import "time"

func main() {
  go helloWorld()

  time.Sleep(time.Second)
}

func helloWorld(){
  println("Hello World!")
}
