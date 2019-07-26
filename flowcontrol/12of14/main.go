package main

import "fmt"

func addNumber(number1 int) int {
	number1 = 8
	fmt.Println(number1)
	return number1
}

func main() {
	number := 10

	defer fmt.Println(addNumber(number))

	fmt.Println(number)
}
