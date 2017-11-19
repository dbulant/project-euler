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
	return 1
}

func isNumberPalindrome(n uint64) bool {
	d := getDigitsFromNumber(n)
	for i, j := 0, len(d); i < len(d)/2; i++ {
		if d[i] != d[j-i-1] {
			return false
		}
	}
	return true
}

func getDigitsFromNumber(n uint64) []uint64 {
	digits := make([]uint64, 0, 10)
	var base uint64 = 10
	for n > 0 {
		r := n % base
		n -= r
		n /= base
		rs := make([]uint64, 0, 1)
		rs = append(rs, r)
		digits = append(rs, digits...)
		rs = nil
	}
	return digits
}
