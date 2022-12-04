package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	min := 10000000000
	list := myutil.InputNumList()

	for _, v := range list {
		if v < min {
			min = v
		}
	}
	fmt.Println(min)
}
