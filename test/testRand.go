package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		a := rand.Intn(100)
		b := rand.Intn(100)
		fmt.Printf("a=%d\tb=%d\n", a, b)
	}
}
