package main

import (
	"fmt"
	"math"
)

var a, b, c []int
var aDp, bDp, cDp []int

func main() {
	var n int
	fmt.Scan(&n)
	aDp = make([]int, n)
	bDp = make([]int, n)
	cDp = make([]int, n)
	for i := 0; i < n; i++ {
		var ai, bi, ci int
		fmt.Scan(&ai, &bi, &ci)
		a = append(a, ai)
		b = append(b, bi)
		c = append(c, ci)
	}
	res := int(math.Max(float64(maxWork(0, n-1)), float64(maxWork(1, n-1))))
	res = int(math.Max(float64(res), float64(maxWork(2, n-1))))
	fmt.Println(res)
}

func maxWork(work, i int) int {
	var val int
	var valMax int
	switch work {
	case 0:
		if aDp[i] != 0 {
			return aDp[i]
		}
		if i == 0 {
			aDp[i] = a[i]
			return aDp[i]
		}
		if maxWork(1, i-1) < maxWork(2, i-1) {
			val = maxWork(2, i-1)
		} else {
			val = maxWork(1, i-1)
		}
		aDp[i] = val + a[i]
		valMax = aDp[i]
	case 1:
		if bDp[i] != 0 {
			return bDp[i]
		}
		if i == 0 {
			bDp[i] = b[i]
			return bDp[i]
		}
		if maxWork(0, i-1) < maxWork(2, i-1) {
			val = maxWork(2, i-1)
		} else {
			val = maxWork(0, i-1)
		}
		bDp[i] = val + b[i]
		valMax = bDp[i]
	case 2:
		if cDp[i] != 0 {
			return cDp[i]
		}
		if i == 0 {
			cDp[i] = c[i]
			return cDp[i]
		}
		if maxWork(0, i-1) < maxWork(1, i-1) {
			val = maxWork(1, i-1)
		} else {
			val = maxWork(0, i-1)
		}
		cDp[i] = val + c[i]
		valMax = cDp[i]
	}
	return valMax
}
