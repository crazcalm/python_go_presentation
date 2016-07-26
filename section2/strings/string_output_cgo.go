package main

import "C"
import "bufio"
import "fmt"
import "os"

//export string_output
func string_output() *C.char {
	fmt.Printf("Please enter a string: ")
	input := bufio.NewScanner(os.Stdin)
	var result string

	for input.Scan() {
		if input.Text() != "" {
			result = input.Text()
			break
		} else {
			fmt.Printf("Please enter a string.\nTry again: ")
		}
	}
	cstr := C.CString(result)

	return cstr
}

func main() {
	fmt.Println(string_output())
}
