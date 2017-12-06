// problem14
package projecteuler

import "sync"

//Longest Collatz sequence
//The following iterative sequence is defined for the set of positive integers:
//n → n/2 (n is even)
//n → 3n + 1 (n is odd)
//Using the rule above and starting with 13, we generate the following sequence:
//13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
//It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
//Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
//Which starting number, under one million, produces the longest chain?
//NOTE: Once the chain starts the terms are allowed to go above one million.
//TODO add more go routines, var numCPU = runtime.GOMAXPROCS(0),
//TODO get rid of 4 and hardcoded values,
//TODO create general function that computes something in interval
func Problem14() uint64 {
	number, _ := GeneralProblem14(1000000)
	return number
}

func GeneralProblem14(max uint64) (uint64, []uint64) {
	c := make(chan CollatzResults, 4)
	var wg sync.WaitGroup

	GeneralProblem14InInterval(max/4, 0, &wg, c)
	GeneralProblem14InInterval(max/2, max/4, &wg, c)
	GeneralProblem14InInterval(3*max/4, max/2, &wg, c)
	GeneralProblem14InInterval(max, 3*max/4, &wg, c)
	wg.Wait()

	results := make([]CollatzResults, 0, 4)
	results = append(results, <-c)
	results = append(results, <-c)
	results = append(results, <-c)
	results = append(results, <-c)

	mr := results[0]
	lch := len(results[0].chain)
	for _, cr := range results {
		if len(cr.chain) > lch {
			mr = cr
		}
	}

	return mr.number, mr.chain
}

type CollatzResults struct {
	chain  []uint64
	number uint64
}

func GeneralProblem14InInterval(top uint64, bottom uint64, wg *sync.WaitGroup, c chan CollatzResults) {
	wg.Add(1)
	go func() {
		var currentResult CollatzResults
		for i := top; i > bottom; i-- {
			chain := ComputeCollatzSequence(i)
			if len(chain) > len(currentResult.chain) {
				currentResult.number = i
				currentResult.chain = chain
			}
		}
		c <- currentResult
		wg.Done()
	}()
}

func ComputeCollatzSequence(startingnumber uint64) []uint64 {
	chain := make([]uint64, 0, 10)
	for startingnumber != 1 {
		chain = append(chain, startingnumber)
		startingnumber = Iterate(startingnumber)
	}

	return chain
}

func EvenSequence(n uint64) uint64 {
	return n / 2
}

func OddSequence(n uint64) uint64 {
	return 3*n + 1
}

func Iterate(n uint64) uint64 {
	if n%2 == 0 {
		return EvenSequence(n)
	} else {
		return OddSequence(n)
	}
}
