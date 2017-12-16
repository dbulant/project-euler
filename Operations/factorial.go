// factorial
package projecteuler

func FactorialUint64(n uint64) uint64 {
	var product uint64 = 1
	for i := uint64(1); i <= n; i++ {
		product *= i
	}
	return product
}

func FactorialFloat64(n float64) float64 {
	var product float64 = 1
	for i := float64(1); i <= n; i++ {
		product *= i
	}
	return product
}

//Computes factorial in interval <low,top>
func FactorialInInterval(low uint64, top uint64) uint64 {
	var product uint64 = 1
	var i uint64 = low
	for i <= top {
		product *= i
		i++
	}
	return product
}
