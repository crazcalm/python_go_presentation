package main

import "C"
import "math/rand"
import "time"
import "fmt"

//export string_output
func string_output() string {
    // Look into seeding a random generator
    // figure out how to get random number.
    rand.Seed(time.Now().UTC().UnixNano())
    length := rand.Intn(10)

    bytes := make([]byte, 1)
    for i :=0; i < length; i++ {
        bytes[i] = byte(randInt(65, 90))
    }
    return string(bytes)
}

func randInt(min int, max int) int {
    return min + rand.Intn(max-min)
}

func main(){
    fmt.Println(string_output())
}
