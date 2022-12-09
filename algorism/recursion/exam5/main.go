package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	val := myutil.Input()
	list := myutil.InputNumList()

	fmt.Println(sumCheck(len(list), val, list))
}

func sumCheck(i, val int, list []int) bool {

	if i == 0 {
		if val == 0 {
			return true
		} else {
			return false
		}
	}
	if sumCheck(i-1, val, list) {
		return true
	}
	if sumCheck(i-1, val-list[i-1], list) {
		return true
	}
	return false
}
