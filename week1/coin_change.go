package main

import "fmt"

func main() {
	coins := []int{1, 2, 5}
	amount := 100

	dp := make(map[int]int, amount+1)

	fmt.Println(coinChange(coins, amount, dp))
}

func coinChange(coins []int, amount int, dp map[int]int) int {

	result := coinCh(coins, amount, dp)
	if result >= 1e5 {
		return -1
	}
	return result
}

func coinCh(coins []int, amount int, dp map[int]int) int {

	if amount < 0 {
		// 10 to the power 5 -> no path
		return 1e5
	} else if amount == 0 {
		return 0
	}

	smallest := int(1e5)
	for _, coin := range coins {
		if _, ok := dp[amount-coin]; !ok {
			dp[amount-coin] = coinCh(coins, amount-coin, dp)
		}
		if dp[amount-coin] < smallest {
			smallest = dp[amount-coin]
		}
	}

	// choosing the smallest count of coins
	return 1 + smallest
}
