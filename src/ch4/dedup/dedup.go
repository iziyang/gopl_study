package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {  // 使用 map 来实现集合功能，将元素作为键，值为 true 或 false
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
		fmt.Println(seen)
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup:%v\n", err)
		os.Exit(1)
	}

}
