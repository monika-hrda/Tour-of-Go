package main

import "fmt"

func adder() func(int) int {
	sum := 0
	fmt.Println("value of sum is ", sum)
	return func(x int) int {
		fmt.Println("value of sum inside of the returned function is ", sum)
		sum += x
		return sum
	}
}

func main() {
	pos := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i))
	}
}
