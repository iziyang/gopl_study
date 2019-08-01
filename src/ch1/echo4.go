package main

import (
	"flag"
	"fmt"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	p := new(int) // 初始化一个地址
	fmt.Println(p) //输出地址
	fmt.Println(*p) // 输出值
	*p = 2
	fmt.Println(p)

}
