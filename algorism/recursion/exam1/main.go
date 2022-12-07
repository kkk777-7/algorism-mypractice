package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	n := myutil.Input()
	calc(n)
}

func calc(n int) int {
	fmt.Printf("func(%d)を呼び出しました\n", n)

	if n == 0 {
		return 0
	}

	result := n + calc(n-1)
	fmt.Printf("%dまでの和=%d\n", n, result)
	return result
}
