package main

import (
	"fmt"
	"os"
)

func sum(vals ...int) int { // 变长参数7
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func errorf(linenum int, format string, args ...interface{}) { // interface 类型代表着这个函数最后一个参数可以接受任何值
	fmt.Fprintf(os.Stderr, "Line %d:", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func main() {
	values := []int{1, 2, 3, 4}
	fmt.Println(sum())
	fmt.Println(sum(1))
	fmt.Println(sum(2))
	fmt.Println(sum(3))
	fmt.Println(sum(values...))

	linenum, name := 12, "count"
	errorf(linenum, "undefined: %s", name)
}
