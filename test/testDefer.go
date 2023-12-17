package main

import "fmt"

// 规则一 当defer被声明时，其参数就会被实时解析
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// 规则二 defer执行顺序为先进后出
func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

// 规则三 defer可以读取有名返回值
func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	a()
	b()
	fmt.Println()
	fmt.Println(c())
}
