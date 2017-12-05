// problem20
package projecteuler

func Problem20() {

}

func GeneralProblem20(n uint64) uint64 {
	product := factorial(n)
	digits := getDigitsFromNumber(product)
	sum := sumOfSlice(digits)
	return sum
}

func factorial(n uint64) uint64 {
	var product uint64 = 1
	for i := uint64(1); i <= n; i++ {
		product *= i
	}
	return product
}
