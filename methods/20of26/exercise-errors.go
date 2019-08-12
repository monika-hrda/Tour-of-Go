package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

/* call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop: 
if the Error() method calls fmt.Sprint(e), 
fmt.Sprint(e) will call e.Error() to convert the value of e to a string,
then the program will recurse until out of memory
*/
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := float64(1)
	old := z
	for {
		old = z
		z -= (z*z - x) / (2 * z)

		diff := (z - old)
		if math.Abs(diff) < 0.000000000000001 {
			break
		}
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(15))
	fmt.Println(Sqrt(-2))
}

