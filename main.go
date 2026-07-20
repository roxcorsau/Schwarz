package main

import "fmt"

func TwoSum(nums []int, target int) map[int]int {
	pairs := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				pairs[i] = j
			}
		}
	}

	return pairs
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	result := TwoSum(numbers, target)
	fmt.Printf("Result: %v\n", result)
}
