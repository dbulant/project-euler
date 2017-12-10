// problem36
package projecteuler

import . "github.com/dbulant/project-euler/Operations"

/*
The decimal number, 585 = 1001001001 in binary, is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros.)
*/
func Problem36() uint64 {
	var i uint64 = 2
	var sum uint64 = 0
	for i < 1000000 {
		if IsNumberPalindrome(i) && IsBinaryNumberPalindrome(i) {
			sum += i
		}
		i++
	}
	return sum
}
