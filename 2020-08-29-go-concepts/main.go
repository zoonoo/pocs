package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Hello, 세계")
	io()
}

func cli() {
	var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func io() {
	fmt.Println("io run.")
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		counts[s]++
	}
	for line, n := range counts {
		if n >= 1 {
			fmt.Printf("%d, %T\n", n, line)
		}
	}
}
