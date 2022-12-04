package myutil

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Input() int {
	var n int
	fmt.Scan(&n)
	return n
}

func InputNumList() []int {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var nums []int
	n := nextInt(sc)
	for i := 0; i < n; i++ {
		nums = append(nums, nextInt(sc))
	}
	return nums
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		log.Fatal(err)
	}
	return i
}
