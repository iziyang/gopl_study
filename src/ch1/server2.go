package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main()  {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

// 处理程序回显请求的 URL 的路径部分
func handler2(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

}

//counter 回显目前为止调用的次数
func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()

}
