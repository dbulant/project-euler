// problem4
package projecteuler

import "math"

//A palindromic number reads the same both ways.
//The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
//Find the largest palindrome made from the product of two 3-digit numbers.
func Problem4() uint64 {
	r := GeneralProblem4(3)
	return r
}

func GeneralProblem4(dignumber int) uint64 {
	max := uint64(math.Pow10(dignumber))
	for i := max - 1; i >= 1; i-- {
		for j := max - 1; j >= 1; j-- {
			cur := i * j
			if isNumberPalindrome(cur) {
				return cur
			}
		}
	}
	return 0
}
