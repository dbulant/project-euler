// problem10
package projecteuler

func Problem10() uint64 {
	sum := GeneralProblem10(2000000)
	return sum
}

func GeneralProblem10(max uint64) uint64 {
	sum := sumAllPrimesBelowMaxValue(max)
	return sum

}
