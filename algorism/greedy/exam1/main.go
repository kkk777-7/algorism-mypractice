package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	list := myutil.InputNumList()
	sum := myutil.Input()
	var chk int
	for _, v := range list {
		rest, cnt := sub(sum, v)
		chk += cnt
		sum = rest
	}
	fmt.Println(chk)
}

func sub(sum, n int) (int, int) {
	cnt := sum / n
	return sum - n*cnt, cnt
}
