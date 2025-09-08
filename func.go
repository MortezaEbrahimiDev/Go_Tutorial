package main

import "fmt"

func main() {

	SayHello("morteza")
	SayHelloFullname("Morteza", "Ebrahimi")
	fmt.Printf("sum 5,6 is :%d", Sum(5, 6))
}

func SayHello(name string) {
	fmt.Printf("Hello %s", name)
}

func SayHelloFullname(firstName, LastName string) {
	fmt.Printf("Hello %s %s", firstName, LastName)
}

func Sum(a, b int) int {
	return a + b
}
