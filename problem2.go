// problem2
package main

func Problem2() int {
	sum := GeneralProblem2(4000000)
	return sum
}

func GeneralProblem2(top int) int {
	sum, fn, fn1, fn2 := 0, 0, 1, 1
	for fn < top {
		fn = fn1 + fn2
		fn1 = fn2
		fn2 = fn
		if fn%2 == 0 {
			sum += fn
		}
	}
	return sum
}
