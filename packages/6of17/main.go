package main

import "fmt"

func swapping(a, b string) (string, string) {
	return b, a
}
func main() {
	a, b := swapping("hello", "world")
	fmt.Println(a, b)
}
