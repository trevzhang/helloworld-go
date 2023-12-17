package main

import "fmt"

func main() {
	testSlicePointer1()
}

// 切片是指针类型，修改切片会修改数组
func testSlicePointer1() {
	a1 := [3]int{1, 2, 3}
	fmt.Println("array: \t\t", a1)
	s1 := a1[1:3]
	fmt.Println("slice: \t\t", s1)
	s1[0] = 20
	s1[1] = 30
	fmt.Println("modified array: ", a1)
}

// 切片扩容后修改切片，不会影响原有数组
func testSlicePointer2() {
	a1 := [3]int{1, 2, 3}
	fmt.Println("array: \t\t", a1)
	s1 := a1[1:3]
	fmt.Println("slice: \t\t", s1)
	s1 = append(s1, 4, 5, 6)
	s1[0] = 20
	s1[1] = 30
	fmt.Println("modified slice: ", s1)
	fmt.Println("modified array: ", a1)
}
