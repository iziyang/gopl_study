package main

import (
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}

}

func PopCount(x uint64) int {
	fmt.Println(pc[byte(x>>(0*8))])
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func forPopCount(x uint64) int {
	var s byte
	var i int
	for i = 0; i < 8; i++ {
		//s = s + pc[byte(x>>(i*8))]
	}
	return int(s)
}

var x uint64 = 24

func main() {
	fmt.Println(x)
	fmt.Println(PopCount(x))
	//fmt.Println(forPopCount(x))

}
