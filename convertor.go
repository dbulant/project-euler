// convertor
package projecteuler

import "fmt"

//number to binary
//5 101
//4 10
//8 1000
func uint64ToBinary(number uint64) []uint64 {
	digs := make([]uint64, 0, 10)
	var base uint64 = 2
	if number%2 == 0 {
		order := 0
		for number >= 1 {
			number /= base
			order++
			digs = append(digs, 0)
		}
		digs[0] = 1
	} else {
		//we know that most right number in digs is 1
		number -= 1
		digs = append(digs, 0)
		order := 0
		//find nearest n*n
		for number >= 1 {
			number /= base
			order++
			digs = append(digs, 0)
			digs[order] = 1
		}
		digs[len(digs)-1] = 1

	}

	return digs
}

//work only for pure n^2, not composed, one zero missing or redundant
func evenUint64ToBinary(number uint64) []uint64 {
	digs := make([]uint64, 1, 10)
	var base uint64 = 2
	//var order uint64 = 0
	for number > 1 {
		for number%2 == 0 {
			fmt.Printf("number is %d \n", number)
			number /= base
			if number%2 != 0 && number != 1 {
				number -= 1
				fmt.Printf("number first if is %d \n", number)
				digs = append(digs, 1)
			} else {
				digs = append(digs, 0)
				fmt.Printf("number first else is %d \n", number)
			}
			fmt.Printf("end of one for \n")
		}
	}
	digs[0] = 1
	return digs
}
