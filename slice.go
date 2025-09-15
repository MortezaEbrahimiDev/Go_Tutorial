package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(nums); i++ {
		nums = append(nums, i*i)
	}

	for h := 0; h < len(nums); h++ {
		fmt.Println("Number %d=%d", h, nums[h])
	}
}
