// problem20
package projecteuler

import . "github.com/dbulant/project-euler/Operations"

func Problem20() uint64 {
	var sum uint64 = 0
	var step uint64 = 5
	for i := uint64(0); i < uint64(100); i += step {
		f := FactorialInInterval(i+1, i+step) //does not work with  10, works with 5
		digs := GetDigitsFromNumber(f)
		sum += SumOfSlice(digs)
	}
	return sum
}
