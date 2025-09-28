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
		fmt.Print("عدد اول را وارد کنید (یا 'exit' برای خروج): ")
		input1, _ := reader.ReadString('\n')
		input1 = strings.TrimSpace(input1)
		if strings.ToLower(input1) == "exit" {
			fmt.Println("خروج از برنامه. خداحافظ!")
			break
		}

		num1, err := strconv.ParseFloat(input1, 64)
		if err != nil {
			fmt.Println("لطفا یک عدد معتبر وارد کنید!")
			continue
		}

		fmt.Print("عدد دوم را وارد کنید (یا 'exit' برای خروج): ")
		input2, _ := reader.ReadString('\n')
		input2 = strings.TrimSpace(input2)
		if strings.ToLower(input2) == "exit" {
			fmt.Println("خروج از برنامه. خداحافظ!")
			break
		}

		num2, err := strconv.ParseFloat(input2, 64)
		if err != nil {
			fmt.Println("لطفا یک عدد معتبر وارد کنید!")
			continue
		}

		result := multiply(num1, num2)
		fmt.Printf("نتیجه ضرب: %.2f\n\n", result)
	}
}

func multiply(a, b float64) float64 {
	return a * b
}
