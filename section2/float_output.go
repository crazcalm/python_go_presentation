package main

import "C"
import "math/rand"
import "time"

func init(){
    rand.Seed(time.Now().UnixNano())
}

//export float_output
func float_output() float32 {
    // Look into seeding a random generator
    // figure out how to get random number.
    return (rand.Float32() + float32(1)) * float32(rand.Intn(1000))
}

func main(){}
