package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	v := myutil.Input()
	nums := myutil.InputNumList(7)
	found_id := -1
	for i := 0; i < 7; i++ {
		if nums[i] == v {
			found_id = i
		}
	}
	fmt.Println(found_id)
}
