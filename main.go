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

func TwoSumFast(nums []int, target int) map[int]int {
	pairs := make(map[int]int)
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if prevIdx, found := seen[complement]; found {
			pairs[prevIdx] = i
		}
		seen[num] = i
	}

	return pairs
}

func main() {
	numbers := []int{2, 7, 11, 15, 3, 6}
	target := 9
	fmt.Println("TwoSum (O(n^2)):", TwoSum(numbers, target))
	fmt.Println("TwoSumFast (O(n)):", TwoSumFast(numbers, target))
}
