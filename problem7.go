// problem7
package projecteuler

//By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
//What is the 10 001st prime number?
func Problem7() uint64 {
	pr := GeneralProblem7(10001)
	return pr
}

func GeneralProblem7(amount uint64) uint64 {
	//1 is not prime number, 2 and 3 are
	var found uint64 = 2
	var i uint64 = 3
	for found != amount {
		i++
		if isNumberPrime(i) {
			found++
		}
	}
	return i
}
