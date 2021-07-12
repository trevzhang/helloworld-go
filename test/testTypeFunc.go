package main

import "fmt"

// 测试自定义方法类型
type myFunc func(int) int

func (f myFunc) sum(a, b int) int {
	res := a + b
	return f(res)
}

func mul10(num int) int {
	return num * 10
}

func mul100(num int) int {
	return num * 100
}

func handlerSum(handler myFunc, a, b int) int {
	res := handler.sum(a, b)
	fmt.Println(res)
	return res
}

func main() {
	newFunc1 := myFunc(mul10)
	newFunc2 := myFunc(mul100)

	handlerSum(newFunc1, 1, 1) // 20
	handlerSum(newFunc2, 1, 1) // 200
}
