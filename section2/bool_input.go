package main

import "C"
import "fmt"

//export bool_input
func bool_input(b bool) {
    fmt.Println(b)
}

func main(){}

