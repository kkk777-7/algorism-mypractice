package myutil

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Input() int {
	var n int
	fmt.Scan(&n)
	return n
}

func InputNumList() []int {
	sc.Split(bufio.ScanWords)
	var nums []int
	n := nextInt()
	for i := 0; i < n; i++ {
		nums = append(nums, nextInt())
	}
	return nums
}

func nextInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		log.Fatal(err)
	}
	return i
}
