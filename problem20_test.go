// problem20_test
package projecteuler

import "testing"

func TestFactorial(t *testing.T) {
	var expected float64 = 3628800
	f := factorial(10)
	if f != expected {
		t.Errorf("expected %e, got %e instead \n", expected, f)
	}

	//TODO no hardcoded values this way
	expected = 9.33262154439441e+157
	f = factorial(100)
	//TODO obviously do not use this operator
	if f != expected {
		t.Errorf("expected %g \n", expected)
		t.Errorf("got %g instead \n", f)
		t.Errorf("difference is %g \n", expected-f)
	}

	data, err := getDigitsFromFloatNumber(f)
	if err != nil {
		t.Errorf("Errot during string-float operstions is %v \n", err)
	}
	t.Errorf("data struct is %+v \n", data)
}

func TestGeneralProblem20(t *testing.T) {
	var expected uint64 = 27
	r := GeneralProblem20(10)
	if r != expected {
		t.Errorf("expected %d, got %d instead \n", expected, r)
	}

	r = Problem20()
	if r != 27 {
		t.Errorf("expected %d, got %d instead \n", 27, r)
	}
}
