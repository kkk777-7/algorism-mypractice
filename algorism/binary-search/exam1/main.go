package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	age := myutil.Input()
	list := myutil.InputNumList()
	chk, idx := binarySearch(age, list)
	fmt.Println(chk, idx)
}

func binarySearch(age int, list []int) (bool, int) {
	left := 0
	right := len(list) - 1
	for left <= right {
		mid := left + (right-left)/2
		if list[mid] == age {
			return true, mid
		} else if list[mid] < age {
			left = mid + 1
		} else {
			left = mid - 1
		}
		fmt.Println(left, right, mid, list[mid])
	}
	return false, -1
}
