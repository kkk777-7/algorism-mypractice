package main

import "github.com/kkk777-7/algorism-mypractice/myutil"

func main() {
	list := myutil.InputNumList()

}

func sortBinary(list []int) []int {

	if len(list) == 1 {
		return list
	}
	key := list[0]
	big := make([]int, len(list)/2)
	small := make([]int, len(list)/2)
	for i := 1; i < len(list); i++ {
		if list[i] > key {
			big = append(big, list[i])
		} else {
			small = append(small, list[i])
		}
	}
	list[len(list)/2] = key
}
