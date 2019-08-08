package echo3

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Println(os.Args[1:])

	// 打印索引和元素
	for nu, arg := range os.Args[1:] {
		fmt.Println(nu, arg)
	}

}
