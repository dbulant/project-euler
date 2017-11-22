// problem1_test
package projecteuler

import "testing"

func TestGeneralProblem1(t *testing.T) {
	s := GeneralProblem1(10, []int{3, 5})
	if s != 23 {
		t.Errorf("Expected result is: %d, got: %d \n", 23, s)
	}
}

func TestGetDigitsFromNumber(t *testing.T) {
	d := getDigitsFromNumber(123)
	if d[0] != 1 && d[1] != 2 && d[2] != 3 {
		rd := []uint64{1, 2, 3}
		t.Errorf("digits are %+v \n", d)
		t.Errorf("expected result is %+v \n", rd)
	}

	d = getDigitsFromNumber(654)
	if d[0] != 6 && d[1] != 5 && d[2] != 4 {
		rd := []uint64{6, 5, 4}
		t.Errorf("digits are %+v \n", d)
		t.Errorf("expected result is %+v \n", rd)
	}
}

func TestIsNumberPalindrome(t *testing.T) {
	var tn uint64 = 12345
	r := isNumberPalindrome(tn)
	if r {
		t.Errorf("number %d is not palindrome, yet tests like true \n", tn)
	}

	tn = 12345
	r = isNumberPalindrome(tn)
	if r {
		t.Errorf("number %d is not palindrome, yet tests like true \n", tn)
	}

	tn = 12221
	r = isNumberPalindrome(tn)
	if !r {
		t.Errorf("number %d is palindrome, yet tests like false \n", tn)
	}

	tn = 32123
	r = isNumberPalindrome(tn)
	if !r {
		t.Errorf("number %d is palindrome, yet tests like false \n", tn)
	}

}

func TestGeneralProblem4(t *testing.T) {
	n := GeneralProblem4(2)
	var expected uint64 = 9009
	if n != expected {
		t.Errorf("expected result is %d, got instead %d \n", expected, n)
	}
}

func TestGeneralProblem7(t *testing.T) {
	pn := GeneralProblem7(10001)
	if pn != 104743 {
		t.Errorf("expected result is %d, got instead %d \n", 104743, pn)
	}
}

func TestGeneralProblem10(t *testing.T) {
	sum := Problem10()
	if sum != 142913828922 {
		t.Errorf("expected result is %d, got instead %d \n", 142913828922, sum)
	}
}

func TestProblem13(t *testing.T) {
	//s := Problem13()
	//t.Errorf("first ten digits are: %+v \n", s)
}
