package main

import (
	"fmt"
	"time"
)

//func main() {
//	go printMessage("Hello", 5)
//	go printMessage("world", 5)
//	time.Sleep(1 * time.Second)
//}

func printMessage(msg string, count int) {
	for i := 0; i < count; i++ {
		fmt.Println(msg)
		time.Sleep(200 * time.Millisecond)
	}
}
