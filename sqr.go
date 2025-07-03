package main 

import(
	"fmt"
	"math"
)

func Sqrt(x float64, maxDepth int) float64 {
	fmt.Println("Expected :", math.Sqrt(x))

	if x < 0 {
		return math.NaN()
	}

	z := 1.0
	depth := 0

	for depth= 1; depth  < maxDepth; depth++ {
		prev := z

		z-= (z*z - x) / (2*z)

		if math.Abs(z-prev) < 1e-10 {
			break
		}
	}

	fmt.Println("Ran with depth : ", depth)

	return z
}

func main() {
	fmt.Println(Sqrt(20000, 50))
}
