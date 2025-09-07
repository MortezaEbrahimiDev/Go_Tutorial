package main

import "fmt"

func main() {
	day := "شنبه"

	switch day {
	case "شنبه":
		fmt.Printf("اول هفته")
	case "جمعه":
		fmt.Print("آخر هفته")
	default:
		fmt.Print("وسط هفته")
	}
}
