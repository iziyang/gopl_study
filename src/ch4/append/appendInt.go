package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	z := appendInt(x, 2)
	fmt.Println(z)
	fmt.Println(len(x), cap(x), len(z), cap(z))
	x = append(x, z...)
	fmt.Println(x)

}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z

}
