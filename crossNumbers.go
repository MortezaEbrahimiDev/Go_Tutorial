package main

import "fmt"

func main() {
	var num1, num2 float32

	fmt.Print("insert number 1:")

	_, err := fmt.Scanln(&num1)

	if err != nil {
		fmt.Println("failed to read number 1", err)
		return
	}

	fmt.Println("insert number 2:")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Println("failed to read number 2", err)
		return
	}

	result := Multiply(num1, num2)
	fmt.Printf("%.2f * %.2f =%.2f", num1, num2, result)
}

func Multiply(num1, num2 float32) float32 {
	return num1 * num2
}
