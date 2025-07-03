package main

import (
	"fmt"
	"math/cmplx"
)

var (
	SomeRune rune = '🧙'
	MaxInt uint64 = 1 << 64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value : %v\n", SomeRune, SomeRune)
	fmt.Printf("Type: %T Value : %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value : %v\n", z, z)

	a := 2
	b := float64(a) + 1.5
	fmt.Printf("Type: %T Value : %v\n", b, b)
}
