package main

import "fmt"

func add(x, y int, z string) int {
	fmt.Println(z)
	return x + y
}

func main() {
	fmt.Println("2 + 4 =", add(2,4, "Computing"))
}
