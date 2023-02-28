package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	nums := myutil.InputNumList(6)

	min, max := 1000000, -1
	for _, v := range nums {
		if v < min {
			min = v
		}
		if max < v {
			max = v
		}
	}
	fmt.Println(max - min)
}
