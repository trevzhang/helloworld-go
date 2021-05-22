package main

import (
	"fmt"
	"sync"
)

// 测试go语言协程
func A() {
	fmt.Println("A协程启动")
}
func main() {
	var wg sync.WaitGroup
	fmt.Println("主线程启动")
	wg.Add(1)
	go func(i int) {
		defer wg.Done()
		A()
	}(1)
	wg.Wait()
	fmt.Println("主线程执行完了")
}
