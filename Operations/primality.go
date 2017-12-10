// primality
package projecteuler

//checks if positive integer is prime,assume this is called with n > 3
//1 is not prime number,2 and 3 are prime numbers
func IsNumberPrime(n uint64) bool {
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	var i uint64 = 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}

	return true
}

//Find all primes whose value does not exceed max value.
func FindAllPrimesBelowMaxValue(max uint64) []uint64 {
	primes := make([]uint64, 0, 1000)
	primes = append(primes, 2, 3)
	for i := uint64(4); i < max; i++ {
		if IsNumberPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func SumAllPrimesBelowMaxValue(max uint64) uint64 {
	//2 + 3 = 5
	var sum uint64 = 5
	for i := uint64(4); i < max; i++ {
		if IsNumberPrime(i) {
			sum += i
		}
	}
	return sum
}
