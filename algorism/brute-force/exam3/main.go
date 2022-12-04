package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	min := 100000000
	chk := myutil.Input()
	list := myutil.InputNumList()
	list2 := myutil.InputNumList()

	for _, v := range list {
		for _, v2 := range list2 {
			if chk < v+v2 && v+v2 < min {
				min = v + v2
			}
		}
	}
	fmt.Println(min)
}
