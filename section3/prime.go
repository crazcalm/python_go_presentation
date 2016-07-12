package main

import "C"

//export prime
func prime(num int) bool {
	result := true
	for i := num / 2; i > 1; i-- {
		if num%i == 0 {
			result = false
		}
	}
	return result
}

func main() {}
