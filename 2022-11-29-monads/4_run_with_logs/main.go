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

type Fn func(int) numberWithLogs

func runWithLogs(x numberWithLogs, f Fn) numberWithLogs {
	newNumberWithLogs := f(x.result)
	return numberWithLogs{
		result: newNumberWithLogs.result,
		logs:   append(x.logs, newNumberWithLogs.logs...),
	}
}

func square(x int) numberWithLogs {
	return numberWithLogs{
		result: x * x,
		logs: []string{
			fmt.Sprintf("Squared %d", x),
		},
	}
}

func main() {
	res := runWithLogs(wrapWithLogs(2), square)
	fmt.Printf("%x\n", res.result)
	fmt.Printf("%s\n", res.logs)
}
