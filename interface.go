package main

import "fmt"

func main() {

	c := Circle{Radius: 5}
	printArea(c)
}

type Shap interface {
	Area() float64
	Test() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Test() {
	fmt.Println("tewst")
}

func printArea(s Shap) {
	fmt.Printf("Area: %f", s.Area())
	fmt.Printf("test: %f", s.Test())
}
