package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("Result: %v", maxProfit(prices))
}

func maxProfit(prices []int) int {
	smallest := prices[0]
	largest := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > largest {
			largest = prices[i]
			profit = maxP(profit, largest-smallest)
		}

		if prices[i] < smallest {
			smallest = prices[i]
			largest = smallest
		}
	}

	return profit
}

func maxP(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
