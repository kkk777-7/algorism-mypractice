package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	var chk bool
	val := myutil.Input()
	list := myutil.InputNumList()

	for _, v := range list {
		if v == val {
			chk = true
		}
	}
	fmt.Println(chk)
}
