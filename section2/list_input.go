package main

import "C"
import "fmt"

//export list_input
func addOne(list [int]int){
    fmt.Println(list)
}

func main(){}
