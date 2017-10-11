package math

import (
	"fmt"
	"strconv"
)

// IsPalindrome check integer value is palindrom
func IsPalindrome(x int) bool {
	// Special cases:
	// As discussed above, when x < 0, x is not a palindrome.
	// Also if the last digit of the number is 0, in order to be a palindrome,
	// the first digit of the number also needs to be 0.
	// Only 0 satisfy this property.
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		fmt.Println(revertedNumber, x)
		x = x / 10
		fmt.Println("=>", x)
	}

	// When the length is an odd number, we can get rid of the middle digit by revertedNumber/10
	// For example when the input is 12321, at the end of the while loop we get x = 12, revertedNumber = 123,
	// since the middle digit doesn't matter in palidrome(it will always equal to itself), we can simply get rid of it.
	return x == revertedNumber || x == revertedNumber/10
}

// IsPalindrome2 another approach but slower than first function
func IsPalindrome2(x int) bool {
	// Special cases:
	// As discussed above, when x < 0, x is not a palindrome.
	// Also if the last digit of the number is 0, in order to be a palindrome,
	// the first digit of the number also needs to be 0.
	// Only 0 satisfy this property.
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	xstr := strconv.Itoa(x)
	var maxLen = len(xstr)
	var half = maxLen / 2
	// reverse string from middle
	for i := 0; i < half; i = i + 1 {
		if xstr[i] != xstr[maxLen-1-i] {
			return false
		}
	}
	return true
}
