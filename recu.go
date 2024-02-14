package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []string{"1", "2", "345", "-347"}
	l := len(a) - 1
	a, b := a[:l], a[l:]
	fmt.Println(rec(a, b))
}

func rec(lval, rval []string) (int, bool) {
	fmt.Println("enter", lval, rval)
	l := len(lval) - 1

	if l == 1 {
		res, _ := strconv.Atoi(rval[0])
		return res
	}
	if l == 2 {
		return 0, false
	}

	res, _ := strconv.Atoi(rval[0])
	return rec(lval[:l], lval[l:]) + res

}
