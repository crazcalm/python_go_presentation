package main

import "C"
import "math/rand"
import "time"

func init() {
	rand.Seed(time.Now().UnixNano())
}

//export float_output
func float_output() float64 {
	// Look into seeding a random generator
	// figure out how to get random number.
	return (rand.Float64() + float64(1)) * float64(rand.Intn(1000))
}

func main() {}
