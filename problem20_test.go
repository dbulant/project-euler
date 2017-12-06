// problem20_test
package projecteuler

import "testing"

func TestBigFactorial(t *testing.T) {
	var expected float64 = 3628800
	f := factorial(10)
	if f != expected {
		t.Errorf("expected %e, got %e instead \n", expected, f)
	}

	//TODO no hardcoded values this way
	expected = 9.33262154439441e+157
	expectedData, _ := NewFloatNumberData(expected)
	f = factorial(100)
	factorialData, _ := NewFloatNumberData(f)
	if !factorialData.isEqual(&expectedData, 20) {
		t.Errorf("factorialData %+v \n", factorialData)
		t.Errorf("expectedData %+v \n", expectedData)
	}

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
