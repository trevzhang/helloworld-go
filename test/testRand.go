package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// rand生成的是伪随机数序列,必须指定不同的种子才会生成不同的随机数序列
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		a := rand.Intn(100)
		b := rand.Intn(100)
		fmt.Printf("a=%d\tb=%d\n", a, b)
	}
}
