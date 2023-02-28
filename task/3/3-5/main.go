package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	nums := myutil.InputNumList(6)
	min := 10000000
	for _, v := range nums {
		if v%2 != 0 {
			fmt.Println("-1")
			return
		}
		if v < min {
			min = v
		}
	}
	cnt := 0
	for min%2 == 0 {
		min /= 2
		cnt++
	}
	fmt.Println(cnt)
}
