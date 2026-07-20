package main

import (
	"fmt"
)

// TwoSum searches for two numbers in nums that add up to target.
// It returns their indices if found, or an error/nil if no pair exists.
//
// Challenge: Can you optimize this from O(n^2) to O(n) time complexity using a map?
func TwoSum(nums []int, target int) []int {
	// TODO: Implement your solution here
	// Current brute-force implementation:
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9

	result := TwoSum(numbers, target)

	if result != nil {
		fmt.Printf("Success! Indices: %v (Values: %d + %d = %d)\n",
			result, numbers[result[0]], numbers[result[1]], target)
	} else {
		fmt.Println("No pair found matching the target sum.")
	}
}