package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {

	var chk bool
	val := myutil.Input()
	list := myutil.InputNumList()

	for i := 0; i < (1 << len(list)); i++ {
		sum := 0
		for j := 0; j < len(list); j++ {
			if i&(1<<j) != 0 {
				sum += list[j]
			}
		}
		if sum == val {
			chk = true
		}
	}
	fmt.Println(chk)
}
