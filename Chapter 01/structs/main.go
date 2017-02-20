package main

import "fmt"

type Person struct {
	Name    string
	Surname string
	Hobbies []string
	id      string
}

func (p *Person) GetFullName() string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}

func main() {
	p := Person{
		Name:    "Mario",
		Surname: "Castro",
		Hobbies: []string{"cycling", "electronics", "planes"},
		id:      "sa3-223-asd",
	}

	fmt.Printf("%s likes %s, %s and %s\n", p.GetFullName(), p.Hobbies[0], p.Hobbies[1], p.Hobbies[2])
}
