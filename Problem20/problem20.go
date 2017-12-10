// problem20
package projecteuler

import . "github.com/dbulant/project-euler/Operations"

func Problem20() uint64 {
	var sum uint64 = 0
	var step uint64 = 5
	for i := uint64(0); i < uint64(100); i += step {
		f := factorialInInterval(i+1, i+step) //does not work with  10, works with 5
		digs := GetDigitsFromNumber(f)
		sum += SumOfSlice(digs)
	}
	return sum
}

func factorialUint64(n uint64) uint64 {
	var product uint64 = 1
	for i := uint64(1); i <= n; i++ {
		product *= i
	}
	return product
}

func factorialFloat64(n float64) float64 {
	var product float64 = 1
	for i := float64(1); i <= n; i++ {
		product *= i
	}
	return product
}

//Computes factorial in interval <low,top>
func factorialInInterval(low uint64, top uint64) uint64 {
	var product uint64 = 1
	var i uint64 = low
	for i <= top {
		product *= i
		i++
	}
	return product
}
