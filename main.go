package main

import "fmt"

// Funcția ta originală O(n^2)
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

// FUNCȚIA NOUĂ: Varianta optimizată care rulează instant, indiferent de mărimea slice-ului
func TwoSumFast(nums []int, target int) map[int]int {
	pairs := make(map[int]int)
	seen := make(map[int]int) // Ține minte: valoare -> index

	for i, num := range nums {
		complement := target - num
		// Verificăm dacă am văzut deja numărul care lipsește până la target
		if prevIdx, found := seen[complement]; found {
			pairs[prevIdx] = i
		}
		// Salvăm numărul curent în memorie
		seen[num] = i
	}

	return pairs
}

func main() {
	numbers := []int{2, 7, 11, 15, 3, 6}
	target := 9

	resultFast := TwoSumFast(numbers, target)
	fmt.Printf("Rezultat funcție rapidă: %v\n", resultFast)

