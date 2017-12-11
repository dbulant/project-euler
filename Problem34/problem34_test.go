// problem34_test
package projecteuler

import "testing"

func TestProblem34(t *testing.T) {
	//slow
	r := Problem34(100000)
	t.Errorf("r is %d \n", r)
}

func TestfactorialOfDigits(t *testing.T) {
	var expected uint64 = 145
	f := factorialOfDigits(145)
	if f != expected {
		t.Errorf("expected %d, instead got %d \n", expected, f)
	}
}
