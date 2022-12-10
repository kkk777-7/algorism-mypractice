package main

import (
	"fmt"
	"math"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

const INF = 10000000000000000

func main() {
	list := myutil.InputNumList()

	minCost := flogStep(list)
	fmt.Println(minCost)
}

func flogStep(list []int) float64 {
	dp := make([]float64, len(list))
	for i := 0; i < len(dp); i++ {
		dp[i] = INF
	}
	dp[0] = 0
	for i := 1; i < len(list); i++ {
		if i == 1 {
			dp[1] = math.Abs(float64(list[1]) - float64(list[0]))
		} else {
			dp[i] = math.Min(dp[i-1]+math.Abs(float64(list[i])-float64(list[i-1])), dp[i-2]+math.Abs(float64(list[i])-float64(list[i-2])))
		}
	}
	return dp[len(list)-1]
}
