package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	nums := myutil.InputNumList(2)
	cnt := 0
	for x := 0; x < nums[1]+1; x++ {
		for y := 0; y < nums[1]+1; y++ {
			if x+y <= nums[1] {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
