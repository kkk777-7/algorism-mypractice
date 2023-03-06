package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	n := myutil.Input()
	vals := make([][]int, n)

	for i := 0; i < n; i++ {
		vals[i] = make([]int, 3)

		nums := myutil.InputNumList(3)
		vals[i][0], vals[i][1], vals[i][2] = nums[0], nums[1], nums[2]
	}

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 3)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j == k {
					continue
				}
				dp[i+1][k] = max(dp[i+1][k], dp[i][j]+vals[i][k])
			}
		}
	}

	max := 0
	for i := 0; i < 3; i++ {
		if max < dp[n][i] {
			max = dp[n][i]
		}
	}
	fmt.Println(max)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
