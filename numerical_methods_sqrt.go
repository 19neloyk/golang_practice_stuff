import "fmt"

func Sqrt(x float64) float64 {
	var z float64 = x
	for i := 0; i < 10; i ++ {
		z -= (z*z - x) / (2*z)
		fmt.Println(z)
	}
	return z
}
