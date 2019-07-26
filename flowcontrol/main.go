package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	old := z
	for {
		old = z
		z -= (z*z - x) / (2 * z)

		diff := (z - old)

		if diff < 0 {
			diff = diff * -1
		}

		if diff < 0.000000000000001 {
			break
		}

		fmt.Println(old)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
