// problem5_test
package projecteuler

import "testing"

func TestCreateUintSlice(t *testing.T) {
	s := createUintSlice(10)
	var i uint64 = 1
	for _, v := range s {
		if v != i {
			t.Errorf("expected result is %d, got instead %d \n", i, v)
			t.Errorf("Generated structure is %+v", s)
		}
		i++
	}
}

func TestIsNumberDivisibleBySlice(t *testing.T) {
	s := createUintSlice(10)
	var n uint64 = 1
	tr := isNumberDivisibleBySlice(n, s)
	if tr {
		t.Errorf("%d should not be divisible by all numbers from %+v \n", n, s)
	}
}

func TestGeneralProblem5(t *testing.T) {
	n := GeneralProblem5(10)
	var expected uint64 = 2520
	if n != expected {
		t.Errorf("expected result is %d, got instead %d \n", expected, n)
	}
}
