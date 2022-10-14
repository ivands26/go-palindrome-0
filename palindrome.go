package main

import (
	"fmt"
	"strconv"
)

func Palindrome(input int) bool {
	// for check if input 1 digit, its return true
	if input > 0 && input < 10 {
		return true
	}
	// convert integer to string
	new := strconv.Itoa(input)
	// looping to check palindrome or not
	for j := 0; j < len(new)/2; j++ {
		if new[j] != new[len(new)-1-j] {
			return false
		}
	}
	return true
}

func CountPalindrome(n, m int) int {
	var count int
	// count how much palindrome from n to m
	for i := n; i <= m; i++ {
		if Palindrome(i) {
			count++
		}

	}
	return count
}

func main() {
	fmt.Println(CountPalindrome(1, 10))    // 9
	fmt.Println(CountPalindrome(100, 150)) // 5
	fmt.Println(CountPalindrome(99, 100))  // 1
	fmt.Println(CountPalindrome(21, 31))   // 1
}
