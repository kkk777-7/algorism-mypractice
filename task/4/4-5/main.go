package main

import (
	"fmt"

	"github.com/kkk777-7/algorism-mypractice/myutil"
)

func main() {
	var ans int
	num := myutil.Input()
	number753(num, 0, 0, &ans)
	fmt.Println(ans)
}

func number753(n, cur, flag int, cnt *int) {
	if n < cur {
		return
	}

	if flag == 0b111 {
		*cnt++
	}
	number753(n, cur*10+7, flag|0b100, cnt)
	number753(n, cur*10+5, flag|0b010, cnt)
	number753(n, cur*10+3, flag|0b001, cnt)
}
