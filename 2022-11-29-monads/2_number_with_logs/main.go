package main

import "fmt"

type numberWithLogs struct {
	result int
	logs   []string
}

func square(x int) numberWithLogs {
	return numberWithLogs{
		result: x * x,
		logs: []string{
			fmt.Sprintf("Squared %d", x),
		},
	}
}

func addOne(x numberWithLogs) numberWithLogs {
	return numberWithLogs{
		result: x.result + 1,
		logs:   append(x.logs, fmt.Sprintf("addOne from %d", x.result)),
	}
}

func main() {
	x1 := addOne(square(2))
	fmt.Printf("%d\n", x1.result)
	fmt.Printf("%s\n", x1.logs)
}
