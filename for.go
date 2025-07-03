package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x,n); v < lim {
		return v;
	}
	return lim;
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	sum := 1 

	for sum < 1000 {
		sum += sum

		if sum % 2 == 0 {
			fmt.Println(sum)
		}
	}

	fmt.Println(sum)

	fmt.Println(
		pow(3, 2, 10),
		pow (3, 3, 20),
		pow(3, 3, 200),
	)
}
