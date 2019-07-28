package main

import "encoding/json"
import "fmt"

type MyObject struct {
	Number int `json:"number"`
	Word string
}
func main(){
	object := MyObject{5, "Packt"}
	oJson, _ := json.Marshal(object)
	fmt.Printf("%s\n", oJson)
}