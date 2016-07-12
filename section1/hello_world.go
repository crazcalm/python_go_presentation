package main

import "C"
import "fmt"

//export helloWorld
func helloWorld() {
	fmt.Println("Hello World")
}

func main() {
	helloWorld()
}
