// primality
package projecteuler

//checks if positive integer is prime,assume this is called with n > 3
func isNumberPrime(n uint64) bool {
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
