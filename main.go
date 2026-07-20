package main

import "fmt"

// TODO: try to collect all pairs that satisfy the condition
func TwoSum(nums []int, target int) [][]int {
	var result [][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, []int{nums[i], nums[j]})
			}
		}
	}
	return result
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	result := TwoSum(numbers, target)
	fmt.Printf("Result: %v\n", result)
}
