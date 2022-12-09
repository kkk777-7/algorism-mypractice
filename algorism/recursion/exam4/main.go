package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

var cache []int

func main() {
	n := myutil.Input()
	cache = make([]int, n)
	for i := 0; i < len(cache); i++ {
		cache[i] = -1
	}

	fibo(n)

	for _, val := range cache {
		fmt.Println(val)
	}
}

func fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if cache[n-1] != -1 {
		return cache[n-1]
	}
	cache[n-1] = fibo(n-1) + fibo(n-2)
	return cache[n-1]
}
