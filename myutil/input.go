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

func InputNumList(n int) []int {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = nextInt(sc)
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
