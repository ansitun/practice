package main

func countSubstrings(s string) int {

	// loop the string and for every index, check the left and right characters
	// if same, add to the pallindrome count and check the next left and right characters
	// for every index, count the pallindromes including the character itself
	// check for pallindromes when the string length is even, with left being the index
	// and right being index + 1; The time complexity is O(n^2)

	// update - can use 2d DP array to note the start and end of a substring
	count := 0
	for i := 0; i < len(s); i++ {
		// odd pallindrome => index = left = right and spread out
		// even pallindrome => index = left & right = index+1 and spread out
		count += pallindromeCount(s, i, i) + pallindromeCount(s, i, i+1)
	}

	return count
}
func pallindromeCount(s string, left int, right int) int {
	count := 0

	for ; left >= 0 && right < len(s); {
		if string(s[left]) == string(s[right]) {
			count++
			left--
			right++
		} else {
			break
		}
	}

	return count
}
