// problem1
package main

func Problem1() int {
	top := 1000
	multiplies := []int{3, 5}
	sum := GeneralProblem1(top, multiplies)
	return sum
}

func GeneralProblem1(top int, multipliers []int) int {
	sum := 0
	for i := 1; i < top; i++ {
		if isNumberMultiplyOfNumberFromSlice(i, multipliers) {
			sum += i
		}
	}
	return sum
}

func isNumberMultiplyOfNumberFromSlice(number int, numbers []int) bool {
	for _, n := range numbers {
		if number%n == 0 {
			return true
		}
	}
	return false

}
