package main

import "C"
import "math/rand"
import "time"

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//export int_output
func int_output() int {
	// Look into seeding a random generator
	// figure out how to get random number.
	return rand.Intn(10000)
}

func main() {}
