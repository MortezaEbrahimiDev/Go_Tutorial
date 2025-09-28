package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Select action:")

		fmt.Println("*:1")
		fmt.Println("/:2")
		fmt.Println("+:3")
		fmt.Println("-:4")

		choise, _ := reader.ReadString('\n')

		choise = strings.TrimSpace(choise)

		if strings.ToLower(choise) == "exit" {
			fmt.Println("exit from aplication ...")
			break
		}

		number1 := readNumber(reader, "Enter number1:")
		number2 := readNumber(reader, "Enter number2:")

		switch choise {
		case "1":
			fmt.Printf("multiply %.2f ,%.2f =%.2f", number1, number2, number1*number2)
		case "2":
			fmt.Printf("devide %.2f ,%.2f =%.2f", number1, number2, number1/number2)
		case "3":
			fmt.Printf("sum %.2f ,%.2f =%.2f", number1, number2, number1+number2)
		case "4":
			fmt.Printf("minus %.2f ,%.2f =%.2f", number1, number2, number1-number2)
		}
	}

}

func readNumber(reader *bufio.Reader, prompt string) float64 {
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "exit" {
			fmt.Println("exit from app ...")
			os.Exit(0)
		}

		num, err := strconv.ParseFloat(input, 32)

		if err != nil {
			fmt.Println("error in read number")
			continue
		}

		return num
	}
}
