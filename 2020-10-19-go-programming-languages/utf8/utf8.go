package utf8

import "fmt"

func init() {
	for i, r := range "Hello, 세계." {
		fmt.Printf("%d, %q\n", i, r)
	}
}
