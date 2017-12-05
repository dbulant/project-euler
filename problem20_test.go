// problem20_test
package projecteuler

import "testing"

func TestFactorial(t *testing.T) {
	var expected uint64 = 3628800
	f := factorial(10)
	if f != expected {
		t.Errorf("expected %d, got %d instead \n", expected, f)
	}
}

func TestGeneralProblem20(t *testing.T) {
	r := GeneralProblem20(10)
	if r != 27 {
		t.Errorf("expected %d, got %d instead \n", 27, r)
	}
}
