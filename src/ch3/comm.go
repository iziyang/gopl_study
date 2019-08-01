package main

import (
	"bytes"
	"fmt"
)

func comm(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comm(s[:n-3]) + "," + s[n-3:]
}

func commbuff(s string) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}
	for i := 0; i < len(s); {
		buf.WriteString(s[n-3:])
		buf.WriteString(",")
		n-3
	}

}

func main() {
	s := "12345678"
	fmt.Printf("%s\n", comm(s))

}

//package main
//
//import (
//	"fmt"
//	"os"
//)
//
//func main() {
//	for i := 1; i < len(os.Args); i++ {
//		fmt.Printf("  %s\n", comma(os.Args[i]))
//	}
//}
//
////!+
//// comma inserts commas in a non-negative decimal integer string.
//func comma(s string) string {
//	n := len(s)
//	if n <= 3 {
//		return s
//	}
//	return comma(s[:n-3]) + "," + s[n-3:]
//}
