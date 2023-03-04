package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	n := myutil.Input()
	nums := make(map[int]int, n)
	ans := torinacci(n, nums)
	fmt.Println(ans)
}

func torinacci(n int, nums map[int]int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	if _, ok := nums[n]; ok {
		return nums[n]
	}
	num := torinacci(n-1, nums) + torinacci(n-2, nums) + torinacci(n-3, nums)
	nums[n] = num
	return num
}
