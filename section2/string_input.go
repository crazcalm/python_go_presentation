package main

import "C"
import "fmt"

//export string_input
func string_input(name string) {
    fmt.Println(name)
}

func main(){}
