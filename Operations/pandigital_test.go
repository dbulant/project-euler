// pandigital_test
package projecteuler

import "testing"

func TestIsNumberUniqueInSlice(t *testing.T) {
	s := []uint64{2, 2, 2, 2, 2, 2, 2, 2, 2}
	b := isNumberUniqueInSlice(2, s)
	if b {
		t.Errorf("Number should not be unique, slice is %+v \n", s)
	}

	s = []uint64{5, 4, 3, 2, 1, 6, 7, 8, 9}
	b = isNumberUniqueInSlice(1, s)
	if !b {
		t.Errorf("Number should be unique, slice is %+v \n", s)
	}

	s = []uint64{5, 5, 4, 4, 3, 3, 2, 2, 9}
	b = isNumberUniqueInSlice(3, s)
	if b {
		t.Errorf("Number should not be unique, slice is %+v \n", s)
	}
}

func TestIsDecimalDigitBaseUniqueInSlice(t *testing.T) {
	s := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 9, 9}
	b := isDecimalDigitBaseUniqueInSlice(s)
	if b {
		t.Errorf("Number should be pandigital slice is %+v \n", s)
	}

	s = []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b = isDecimalDigitBaseUniqueInSlice(s)
	if !b {
		t.Errorf("Number should be pandigital slice is %+v \n", s)
	}
}
