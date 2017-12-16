// convertor_test
package projecteuler

import (
	"testing"
)

func TestConvertor(t *testing.T) {
	digs := make([]uint64, 0, 10)
	digs = evenUint64ToBinary(14)
	if digs[0] != 1 && digs[1] != 1 && digs[2] != 1 && digs[3] != 0 {
		t.Errorf("binary digits are %+v \n", digs)
	}

	digs = evenUint64ToBinary(32)
	if digs[0] != 1 && digs[1] != 0 && digs[2] != 0 && digs[3] != 0 && digs[4] != 0 && digs[5] != 0 {
		t.Errorf("binary digits are %+v \n", digs)
	}

	digs = oddUint64ToBinary(5)
	if digs[0] != 1 && digs[1] != 0 && digs[2] != 1 {
		t.Errorf("binary digits are %+v \n", digs)
	}

	digs = oddUint64ToBinary(17)
	if digs[0] != 1 && digs[1] != 0 && digs[2] != 0 && digs[3] != 0 && digs[4] != 1 {
		t.Errorf("binary digits are %+v \n", digs)
	}
}
