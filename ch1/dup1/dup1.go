package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("统计输入重复的字符串")
	fmt.Println("请输入字符串(输入quit停止继续输入):")
	for input.Scan() {
		var t string
		if t = input.Text(); t == "quit" {
			break
		}
		counts[t]++
	}
	// 注意:忽略input.Err()中可能的错误
	fmt.Printf("字符串\t出现次数\n")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\n", line, n)
		}
	}
}
