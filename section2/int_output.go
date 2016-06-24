package main

import "C"
import "crypto/rand"

//export int_output
func int_output() int {
    // Look into seeding a random generator
    // figure out how to get random number.
    return rand.Int()
}

func main(){}
