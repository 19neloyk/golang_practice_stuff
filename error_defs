package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	str := fmt.Sprintf("%f", float64(e))
	return "cannot Sqrt negative number: "+ str
}


func Sqrt(x float64) (float64, error) {
	var err error
	if (x < 0) {
		err = ErrNegativeSqrt(x)
		return -1, err
	}
	var z float64 = x
	for i := 0; i < 10; i ++ {
		z -= (z*z - x) / (2*z)
	}
	return z, err
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
