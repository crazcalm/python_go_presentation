package main

import "C"
import "fmt"

//export dict_input
func dict_input(dict map[string]string) {
	for key, value := range dict {
		fmt.Printf("Key is %s -- Value is %s\n", key, value)
	}
}

func main() {}
