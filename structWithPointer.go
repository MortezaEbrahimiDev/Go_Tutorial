package main

import "fmt"

func main() {

	person := Person{Name: "Morteza", Age: 29}
	fmt.Printf("morteza has %d", person.Age)
	person.Birthday()

	fmt.Printf("morteza has %d", person.Age)

}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Birthday() {
	p.Age++
}
