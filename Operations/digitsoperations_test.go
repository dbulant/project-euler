// digitsoperations_test.go
package projecteuler

import (
	"testing"
)

func TestGetDigitsFromNumber(t *testing.T) {
	d := GetDigitsFromNumber(123)
	if d[0] != 1 && d[1] != 2 && d[2] != 3 {
		rd := []uint64{1, 2, 3}
		t.Errorf("digits are %+v \n", d)
		t.Errorf("expected result is %+v \n", rd)
	}

	d = GetDigitsFromNumber(654)
	if d[0] != 6 && d[1] != 5 && d[2] != 4 {
		rd := []uint64{6, 5, 4}
		t.Errorf("digits are %+v \n", d)
		t.Errorf("expected result is %+v \n", rd)
	}
}

func TestIsNumberPalindrome(t *testing.T) {
	var tn uint64 = 12345
	r := IsNumberPalindrome(tn)
	if r {
		t.Errorf("number %d is not palindrome, yet tests like true \n", tn)
	}

	tn = 12345
	r = IsNumberPalindrome(tn)
	if r {
		t.Errorf("number %d is not palindrome, yet tests like true \n", tn)
	}

	tn = 12221
	r = IsNumberPalindrome(tn)
	if !r {
		t.Errorf("number %d is palindrome, yet tests like false \n", tn)
	}

	tn = 32123
	r = IsNumberPalindrome(tn)
	if !r {
		t.Errorf("number %d is palindrome, yet tests like false \n", tn)
	}

}
