package main

import "C"

//export string_output
func string_output() *C.char {
    return C.CString("Hello")
}

func main(){}
