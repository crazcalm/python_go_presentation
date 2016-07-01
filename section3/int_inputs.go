package main

import "C"
import "fmt"

func int_input(nums ...int) {
    for _, num := range nums {
        fmt.Println(num)
    }
}

func main(){
    int_input(1,2,3,4,5,6)
}
