package main

import "C"
import "fmt"

//export list_input
func list_input(list []int) {
	fmt.Println(list)
}

func main() {}
