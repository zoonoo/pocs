package main

import "fmt"

type numberWithLogs struct {
	result int
	logs   []string
}

func wrapWithLogs(x int) numberWithLogs {
	return numberWithLogs{
		result: x,
		logs:   []string{},
	}
}

func square(x numberWithLogs) numberWithLogs {
	return numberWithLogs{
		result: x.result * x.result,
		logs: append(
			x.logs,
			fmt.Sprintf("Squared %d", x.result),
		),
	}
}

func main() {
	fmt.Printf("%d\n", square(square(wrapWithLogs(5))).result)
	fmt.Printf("%s\n", square(square(wrapWithLogs(5))).logs)
}
