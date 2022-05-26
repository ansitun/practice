package main

import "fmt"

// https://leetcode.com/problems/climbing-stairs/
// seems like a fibonacci series
func main() {

	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	memory := make(map[int]int, 0)
	memory[0] = 0
	memory[1] = 1
	memory[2] = 2

	return climbStairsHelper(n, memory)
}

func climbStairsHelper(n int, memory map[int]int) int {
	if _, ok := memory[n-1]; !ok {
		memory[n-1] = climbStairsHelper(n-1, memory)
	}
	if _, ok := memory[n-2]; !ok {
		memory[n-2] = climbStairsHelper(n-2, memory)
	}

	return memory[n-1] + memory[n-2]
}
