// problem20_test
package projecteuler

import "testing"

var EXPECTED_FACTORIAL_100 float64 = 9.33262154439441e+157
var EXPECTED_FACTORIAL_10 float64 = 3628800

func TestFloatNumberData(t *testing.T) {
	//TODO no hardcoded values this way
	expectedData, _ := NewFloatNumberData(EXPECTED_FACTORIAL_100)
	f := factorialFloat64(100)
	factorialData, _ := NewFloatNumberData(f)
	if !factorialData.isEqual(&expectedData, 20) {
		t.Errorf("factorialData %+v \n", factorialData)
		t.Errorf("expectedData %+v \n", expectedData)
	}
}

func TestFactorial(t *testing.T) {
	f := factorialFloat64(10)
	if f != EXPECTED_FACTORIAL_10 {
		t.Errorf("expected %e, got %e instead \n", EXPECTED_FACTORIAL_10, f)
	}
}

func TestBigFactorial(t *testing.T) {
	var final float64 = 1
	var step uint64 = 5
	for i := uint64(0); i < uint64(100); i += step {
		f := factorialInInterval(i+1, i+step) //does not work with  10, works with 5
		final *= float64(f)
	}

	finalData, _ := NewFloatNumberData(final)
	expectedData, _ := NewFloatNumberData(EXPECTED_FACTORIAL_100)
	if !finalData.isEqual(&expectedData, int(step)) {
		t.Errorf("expected %f, got %f instead \n", EXPECTED_FACTORIAL_100, final)
	}
}

func TestFactorialInInterval(t *testing.T) {
	f := float64(factorialInInterval(1, 10))
	if f != EXPECTED_FACTORIAL_10 {
		t.Errorf("expected %f, got %f instead \n", EXPECTED_FACTORIAL_10, f)
	}

	f1 := float64(factorialInInterval(1, 5))
	f2 := float64(factorialInInterval(6, 10))
	if f := f1 * f2; f != EXPECTED_FACTORIAL_10 {
		t.Errorf("expected %f, got %f instead \n", EXPECTED_FACTORIAL_10, f)
	}
}

func TestGeneralProblem20(t *testing.T) {
}
