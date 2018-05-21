package main

import "strconv"

// isPalindrome returns weather an input string is a palendrome.
// the code is copied from the a go playground solution:
// https://play.golang.org/p/5FUOzjBa-o
func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func largestProductPalendrome(maxComponent int) int {
	for n1 := maxComponent; n1 > 0; n1-- {
		for n2 := maxComponent; n2 > 0; n2-- {
			p := n1 * n2

			str := strconv.Itoa(p)

			if isPalindrome(str) {
				return p
			}
		}
	}
	return 1
}

func main() {
	// println(largestProductPalendrome(99)) // example
	println(largestProductPalendrome(999))
}
