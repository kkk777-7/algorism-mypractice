package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	nums := myutil.InputNumList(7)
	min, second := 10000000, 10000000
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	for _, v := range nums {
		if min < v && v < second {
			second = v
		}
	}
	fmt.Println(second)
}
