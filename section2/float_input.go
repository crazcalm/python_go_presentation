package main

import "C"
import "fmt"

//export float_input
func float_input(num float64) {
	fmt.Println(num)
}

func main() {}
