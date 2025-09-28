package main

import "fmt"

func main() {

	result, err := divide(1, 0)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(result)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}

	return a / b, nil
}
