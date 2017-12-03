// problem16
package projecteuler

import "math"

func GeneralProblem16(exponent uint) uint64 {
	var number float64 = math.Pow(2, float64(exponent))
	digits := getDigitsFromNumber(uint64(number))
	sum := sumOfSlice(digits)
	return sum
}
