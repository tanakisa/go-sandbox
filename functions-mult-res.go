package main

import "fmt"

func swap(x, y string) (string,string) {
	return y, x
}

func main () {
	a, b := swap("hello", "world") // declare and assign
	fmt.Println(a,b)

	var c, d string // declare
	c, d = swap("what's", "up") // assign
	fmt.Println(c, d)
}
