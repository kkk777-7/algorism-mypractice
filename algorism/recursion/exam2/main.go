package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	n := myutil.Input()
	m := myutil.Input()
	gcd(n, m)
}

func gcd(n, m int) {
	if m == 0 {
		fmt.Printf("最大公約数は%dです\n", n)
		return
	}
	gcd(m, n%m)
}
