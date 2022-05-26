package main

import "fmt"

// https://leetcode.com/problems/sqrtx/
// determine the square root aof a number
func main() {
	x := 1
	fmt.Printf("Number: %d - Sqrt: %d", x, mySqrt(x))
}

func mySqrt(x int) int {
	return binarySearch(1, x, x)

}

func binarySearch(low int, high int, value int) int {
	if high < low {
		return 0
	}

	mid := (high + low) / 2
	if mid <= value/mid && (mid+1) > value/(mid+1) {
		return mid
	} else if mid > value/mid {
		mid = binarySearch(low, mid-1, value)
	} else {
		mid = binarySearch(mid+1, high, value)
	}

	return mid
}
