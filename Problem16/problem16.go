// problem16
package projecteuler

import "math"
import . "github.com/dbulant/project-euler/Operations"

func GeneralProblem16(exponent uint) uint64 {
	var number float64 = math.Pow(2, float64(exponent))
	digits := GetDigitsFromNumber(uint64(number))
	sum := SumOfSlice(digits)
	return sum
}
