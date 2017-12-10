// problem5
package projecteuler

import . "github.com/dbulant/project-euler/Operations"

//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
func Problem5() uint64 {
	var max uint64 = 20
	n := GeneralProblem5(max)
	return n
}

//Parameter is topDivider, i.e all positive integers from 1 to max
func GeneralProblem5(max uint64) uint64 {
	s := CreateUintSlice(max)
	maxDivider := ProductOfSlice(s)
	var i uint64
	for i = 1; i <= maxDivider; i++ {
		if isNumberDivisibleBySlice(i, s) {
			return i
		}
	}
	return maxDivider
}

func isNumberDivisibleBySlice(n uint64, dividers []uint64) bool {
	for _, v := range dividers {
		if n%v != 0 {
			return false
		}
	}
	return true
}
