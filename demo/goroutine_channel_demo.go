package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 1
	ch := make(chan string, 2)

	// 启动一个goroutine发送消息到通道
	go func() {
		// 使用循环发送消息直到i的值达到1000
		for i <= 1000 {
			ch <- "Hello from channel! " + strconv.Itoa(i)
			i++
		}
		// 关闭通道
		close(ch)
	}()

	// 使用range关键字迭代通道中的消息
	for msg := range ch {
		fmt.Println(msg)
	}
}
