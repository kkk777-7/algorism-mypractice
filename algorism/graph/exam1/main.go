package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	nums := myutil.InputNumList(2)
	n, m := nums[0], nums[1]

	graph := make([][]int, n)

	for i := 0; i < m; i++ {
		vals := myutil.InputNumList(2)
		graph[vals[0]] = append(graph[vals[0]], vals[1])
	}
	fmt.Println(graph)
}
