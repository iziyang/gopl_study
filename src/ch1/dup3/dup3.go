/*
用来检测文件中重复的行，以行作为键，如果有重复的行，那么值+1
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		//fmt.Println(filename)
		data, err := ioutil.ReadFile(filename)
		//fmt.Println(data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}

	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
	//fmt.Println(counts) // map[:3] 空字符串作为键

}
