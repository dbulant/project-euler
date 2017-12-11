// problem34
package projecteuler

import . "github.com/dbulant/project-euler/Operations"

/*
145 is a curious number, as 1! + 4! + 5! = 1 + 24 + 120 = 145.

Find the sum of all numbers which are equal to the sum of the factorial of their digits.

Note: as 1! = 1 and 2! = 2 are not sums they are not included.
*/
func Problem34(top uint64) uint64 {
	var i uint64 = 3
	var sum uint64 = 0
	for i < top {
		digitsSum := factorialOfDigits(i)
		if digitsSum == i {
			sum += digitsSum
		}
		i++
	}

	return sum
}

func factorialOfDigits(n uint64) uint64 {
	digs := GetDigitsFromNumber(n)
	var digitsSum uint64 = 0
	for _, d := range digs {
		f := FactorialUint64(d)
		digitsSum += f
	}
	return digitsSum
}
