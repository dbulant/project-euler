// convertor
package projecteuler

//work only for pure n^2, not composed, one zero missing or redundant
func evenUint64ToBinary(number uint64) []uint64 {
	digs := make([]uint64, 1, 10)
	var base uint64 = 2
	for number > 1 {
		for number%2 == 0 {
			number /= base
			if number%2 != 0 && number != 1 {
				number -= 1
				digs = append(digs, 1)
			} else {
				digs = append(digs, 0)
			}
		}
	}
	digs[0] = 1
	return digs
}

func oddUint64ToBinary(number uint64) []uint64 {
	digs := evenUint64ToBinary(number - 1)
	l := len(digs)
	digs[l-1] = 1
	return digs
}
