// problem20
package projecteuler

func Problem20() uint64 {
	return GeneralProblem20(100)
}

func GeneralProblem20(n uint64) uint64 {
	/*
		product := factorial(n)
		digits := getDigitsFromNumber(product)
		sum := sumOfSlice(digits)
		return sum
	*/
	return 1
}

func factorial(n float64) float64 {
	var product float64 = 1
	for i := float64(1); i <= n; i++ {
		product *= i
	}
	return product
}
