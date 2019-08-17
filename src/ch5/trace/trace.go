package main

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // () 一定要加上去，不然永远不会调用返回的函数变量
	time.Sleep(5 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	bigSlowOperation()
}
