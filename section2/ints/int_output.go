package main

import (
	"C"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//export int_output
func int_output() int {
	fmt.Println("Please enter an integer:")
	input := bufio.NewScanner(os.Stdin)
	var result int

	for input.Scan() {
		text := strings.TrimSpace(input.Text())
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Printf("%s is not a valid integer\nTry again: ", text)
		} else {
			result = num
			break
		}
	}

	return result
}

func main() {
	int_output()
}
