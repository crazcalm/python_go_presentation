package main

import "C"
import "math/rand"
import "time"

//export int_output
func int_output() int {
    // Look into seeding a random generator
    // figure out how to get random number.
    rand.Seed(time.Now().UTC().UnixNano())
    return rand.Intn(10000)
}

func main(){}
