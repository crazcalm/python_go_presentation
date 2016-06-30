package main

import "C"
import "math/rand"
import "time"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//export bool_output
func bool_output() bool {
	// Look into seeding a random generator
	// figure out how to get random number.
	answer := false
	if rand.Intn(2) == 1 {
		answer = true
	}
	return answer
}

func main() {}
