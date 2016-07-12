package main

import "C"

func prime_check(num int) bool {
	result := true
	for i := num / 2; i > 1; i-- {
		if num%i == 0 {
			result = false
		}
	}
	return result
}

func prime(in <-chan int, out chan<- int) {
	for num := range in {
		if prime_check(num) {
			out <- num
		}
	}
	defer close(out)
}

func numList(bound int, out chan<- int) {
	for i := bound; i > 1; i-- {
		out <- i
	}
	defer close(out)
}

func printer(in <-chan int) int {
	count := 0
	for range in {
		count += 1
	}
	return count
}

//export num_of_primes
func num_of_primes(num int) int{
	numbers := make(chan int)
	primes := make(chan int)

	go numList(num, numbers)
	go prime(numbers, primes)
	answer := printer(primes)
	return answer
}

func main(){}
