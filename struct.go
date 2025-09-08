package main

import "fmt"

func main() {

	p := Person{Name: "Morteza", Age: 29}
	p.ShowInfo()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) ShowInfo() {
	fmt.Printf("Person info: Name:%s and age:%d", p.Name, p.Age)
}
