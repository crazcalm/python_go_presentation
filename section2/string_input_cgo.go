package main

import "C"
import "fmt"

//export string_input
func string_input(name *C.char) {
    fmt.Println(C.GoString(name))
}

func main() {}
