package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() { ch1 <- "message 1" }()

	go func() { ch2 <- "massage 2" }()

	select {
	case msg1 := <-ch1:
		fmt.Println("from ch 1:", msg1)

	case msg2 := <-ch2:
		fmt.Println("from ch 2:", msg2)
	}

}
