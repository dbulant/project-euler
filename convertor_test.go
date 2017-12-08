// convertor_test
package projecteuler

import (
	"testing"
)

func TestConvertor(t *testing.T) {
	digs := make([]uint64, 0, 10)
	//digs := evenUint64ToBinary(8)
	//t.Errorf("digs are %+v \n", digs)

	digs = evenUint64ToBinary(32)
	t.Errorf("digs are %+v \n", digs)
	/*
		digs = uint64ToBinary(5)
		t.Errorf("digs are %+v \n", digs)
	*/
}
