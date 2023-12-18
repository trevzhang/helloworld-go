package main

import (
	"fmt"
	"os"
)

func main() {
	str, separate := "", ""
	for _, arg := range os.Args[1:] {
		str += separate + arg
		separate = "-"
	}
	fmt.Println(str)
}
