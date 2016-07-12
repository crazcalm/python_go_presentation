package main

import "C"
import "math/rand"
import "time"

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//export string_output
func string_output() string {
	// Look into seeding a random generator
	// figure out how to get random number.
	length := rand.Intn(20)
	return randomStringRunes(length)
}

func main() {}
