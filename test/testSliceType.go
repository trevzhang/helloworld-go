package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 切片
	s1 := []int{1, 2, 3}
	s2 := make([]int, 3)
	fmt.Println(reflect.TypeOf(s1), s1, reflect.TypeOf(s2), s2)
	// 数组
	a1 := [3]int{1, 2, 3}
	a2 := [...]int{1, 2, 3}
	fmt.Println(reflect.TypeOf(a1), a1, reflect.TypeOf(a2), a2)
}
