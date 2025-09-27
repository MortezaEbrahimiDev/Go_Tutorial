package main

import "fmt"

func main() {
	name := "Morteza"
	go sayHello(name)
	sayHello(name)
}

func sayHello(name string) {

	fmt.Printf("Hello %s", name)
}
