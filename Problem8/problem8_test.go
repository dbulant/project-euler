// problem8_test
package projecteuler

import "testing"

func TestGeneralProblem8(t *testing.T) {
	r := GeneralProblem8(THOUSAND_DIGITS_NUMBER, 4)
	if r != 5832 {
		t.Errorf("expected result is %d, got instead %d \n", 5832, r)
	}
}

func TestProblem8(t *testing.T) {
	r := GeneralProblem8(THOUSAND_DIGITS_NUMBER, 13)
	if r != 23514624000 {
		t.Errorf("expected result is %d, got instead %d \n", 5832, r)
	}
}
