// problem9_test
package projecteuler

import "testing"

func TestIsPairPythagorean(t *testing.T) {
	r := isPairPythagorean(3, 4)
	if !r {
		t.Errorf("expected result is %t, got instead %t \n", true, r)
		t.Errorf("3^2 + 4^2 = 9 + 16 = 25 = 5^2 \n")
	}
}

func TestGeneralProblem9(t *testing.T) {
	arr := GeneralProblem9(1000)
	if arr[0] != 200 || arr[1] != 375 || arr[2] != 425 {
		t.Errorf("result array is %+v \n", arr)
	}
}
