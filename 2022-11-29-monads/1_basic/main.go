package main

import "fmt"

func square(x int) int {
	return x * x
}

func addOne(x int) int {
	return x + 1
}

func main() {
	x1 := addOne(square(2))
	fmt.Printf("%d\n", x1)
}
