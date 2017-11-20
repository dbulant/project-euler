// problem6
package projecteuler

//The sum of the squares of the first ten natural numbers is,
//1^2 + 2^2 + ... + 10^2 = 385
//The square of the sum of the first ten natural numbers is,
//(1 + 2 + ... + 10)^2 = 55^2 = 3025
//Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
//Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
func Problem6() uint64 {
	return GeneralProblem6(100)
}

func GeneralProblem6(max uint64) uint64 {
	sOfSum := squareOfSum(max)
	sOfSquares := sumOfSquares(max)
	return sOfSum - sOfSquares
}

func sumOfSquares(max uint64) uint64 {
	var i, sum uint64 = 1, 0
	for i <= max {
		sum += (i * i)
		i++
	}
	return sum
}

func squareOfSum(max uint64) uint64 {
	var i, sum uint64 = 1, 0
	for i <= max {
		sum += i
		i++
	}
	return sum * sum
}

func sumOfSlice(s []uint64) uint64 {
	var sum uint64 = 0
	for _, v := range s {
		sum += v
	}
	return sum
}
