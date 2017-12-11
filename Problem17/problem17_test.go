// problem17_test.go
package projecteuler

import (
	"testing"
)

func TestGeneralProblem17(t *testing.T) {
	tc := GeneralProblem17(5)
	if tc != 19 {
		t.Errorf("expected result is %d, got instead %d \n", 19, tc)
	}

	ls := NumberToLetters(342)
	lc := CountWithoutHyphensSpaces(ls[0])
	if lc != 23 {
		t.Errorf("expected result is %d, got instead %d \n", 23, lc)
	}

	ls = NumberToLetters(115)
	lc = CountWithoutHyphensSpaces(ls[0])
	if lc != 20 {
		t.Errorf("expected result is %d, got instead %d \n", 20, lc)
	}
}
