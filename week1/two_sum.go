package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Printf("Result: %v", twoSum(nums, 9))
}

func twoSum(nums []int, target int) []int {

	result := make([]int, 0)
	visitedMap := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		if _, ok := visitedMap[target-nums[i]]; ok {
			result = append(result, visitedMap[target-nums[i]], i)

			return result
		} else {
			visitedMap[nums[i]] = i
		}
	}

	return result
}
