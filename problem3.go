// problem3
package projecteuler

import "fmt"

//The prime factors of a positive integer are the prime numbers that divide that integer exactly.
func Problem3() {

}

func GeneralProblem3(primeNumber uint64) {

}

func getAllDividersOfNumber(number uint64) []uint64 {
	dividers := make([]uint64, 0)
	var i uint64
	for i = 2; i < number; i++ {
		if number%i == 0 {
			dividers = append(dividers, i)
		}
	}
	return dividers
}

//Finds all prime numbers less or equal than top value.
//More info: https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes
func SieveOfEratosthenes(top uint64) []uint64 {
	numbers := generateListOfUIntegers(top)
	fmt.Printf("generated numbers: %+v \n", numbers)
	nl := deleteMultipliersFromList(2, numbers)
	numbers = numbers[:nl]
	fmt.Printf("after deleteMultipliers: %+v \n", numbers)
	nl = deleteMultipliersFromList(3, numbers)
	numbers = numbers[:nl]
	fmt.Printf("after deleteMultipliers: %+v \n", numbers)

	nl = deleteMultipliersFromList(5, numbers)
	numbers = numbers[:nl]
	fmt.Printf("after deleteMultipliers: %+v \n", numbers)

	//fmt.Printf("len of numbers is: %d \n", len(numbers))
	//fmt.Printf("cap of numbers is: %d \n", cap(numbers))
	return numbers
}

//Generates list of uint64 less or equal than top value and higher or equal than 2.
func generateListOfUIntegers(top uint64) []uint64 {
	list := make([]uint64, 0, top)
	var i uint64 = 2
	for i <= top {
		list = append(list, i)
		i++
	}
	return list
}

func deleteMultipliersFromList(mul uint64, numbers []uint64) int {
	//fmt.Printf("beginning deleteMultipliers %+v \n", numbers[:])
	j := len(numbers)
	f := make([]uint64, 0, 1)
	f = append(f, numbers[0])
	for i := 0; i < j; i++ {
		if numbers[i]%mul == 0 {
			//fmt.Printf("i is %d \n", i)
			//fmt.Printf("numbers[:i] %+v \n", numbers[:i])
			//fmt.Printf("numbers[i+1:] %+v \n", numbers[i+1:])
			numbers = append(numbers[:i], numbers[i+1:]...)
			j--
		}
	}
	numbers = append(f, numbers...)
	//fmt.Printf("end deleteMultipliers %+v \n", numbers[:])
	return len(numbers)
	//numbers = numbers[:len(numbers)]
	//fmt.Printf("len of numbers is: %d \n", len(numbers))
	//fmt.Printf("cap of numbers is: %d \n", cap(numbers))
}

func testingSSH() {

}
