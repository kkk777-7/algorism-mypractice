package main

import (
	"fmt"
	"math"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

const INF float64 = 10000000000000000

var dp []float64

func main() {
	list := myutil.InputNumList()
	dp = make([]float64, len(list))
	minCost := flogStep(list, len(list)-1)
	fmt.Println(minCost)
}

func flogStep(list []int, i int) float64 {
	// recursion
	if dp[i] < INF {
		return dp[i]
	}

	if i == 0 {
		return 0
	}
	res := INF
	res = math.Min(res, flogStep(list, i-1)+math.Abs(float64(list[i])-float64(list[i-1])))
	if i > 1 {
		res = math.Min(res, flogStep(list, i-2)+math.Abs(float64(list[i])-float64(list[i-2])))
	}
	dp[i] = res
	return dp[i]
}
