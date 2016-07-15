package main

import "C"

//export float_output
func float_output() C.double{
    return C.double(2.4)
}

func main(){}

