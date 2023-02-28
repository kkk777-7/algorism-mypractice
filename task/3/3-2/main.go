package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	val := myutil.Input()
	nums := myutil.InputNumList(5)

	cnt := 0
	for _, v := range nums {
		if v == val {
			cnt++
		}
	}
	fmt.Println(cnt)
}
