package main

import (
	"C"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//export bool_output
func bool_output() bool {
	fmt.Printf("Please enter a boolean value: ")
	input := bufio.NewScanner(os.Stdin)
	var result bool
	for input.Scan() {
		text := strings.TrimSpace(input.Text())

		userBool, err := strconv.ParseBool(text)
		if err != nil {
			fmt.Printf("%s is not a valid answer.\n Try again: ", text)
		} else {
			result = userBool
			break
		}
	}
	return result
}

func main() {
	bool_output()
}
