package main

// #include <stdlib.h>
import "C"
import "unsafe"

//export string_output
func string_output() *C.char {
    var cstr = C.CString("Hello")
    defer C.free(unsafe.Pointer(cstr))
    return cstr
}

func main(){}
