// problem9
package projecteuler

import "math"

//A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
//a^2 + b^2 = c^2
//For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
//There exists exactly one Pythagorean triplet for which a + b + c = 1000.
//Find the product abc.
func Problem9() uint64 {
	arr := GeneralProblem9(1000)
	var product uint64 = 1
	for _, v := range arr {
		product *= v
	}
	return product
}

func GeneralProblem9(n uint64) [3]uint64 {
	var a uint64 = 1
	for a < n {
		var b uint64 = 2
		for b < n {
			if isPairPythagorean(a, b) {
				cs := a*a + b*b
				c := uint64(math.Sqrt(float64(cs)))
				sum := a + b + c
				if sum == n {
					return [3]uint64{a, b, c}
				}
			}
			b++
		}
		a++
	}

	return [3]uint64{0, 0, 0}
}

//Computes a^2 + b^2 and checks if c^(1/2) is integer
func isPairPythagorean(a uint64, b uint64) bool {
	c := a*a + b*b
	cr := math.Sqrt(float64(c))
	if isNumberInteger(cr) {
		return true
	}
	return false
}

func isNumberInteger(x float64) bool {
	return x == math.Floor(x)
}
