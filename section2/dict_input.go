package main

import "C"
import "fmt"

//export dict_input
func addOne(dict map[string]string) {
    fmt.Println(dict)
}

func main(){}
