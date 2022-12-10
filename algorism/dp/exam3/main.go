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
	for i := 0; i < len(list)-1; i++ {
		dp[i+1] = math.Min(dp[i+1], dp[i]+math.Abs(float64(list[i])-float64(list[i+1])))
		if i < len(list)-2 {
			dp[i+2] = math.Min(dp[i+2], dp[i]+math.Abs(float64(list[i])-float64(list[i+2])))
		}
	}
	return dp[len(list)-1]
}
