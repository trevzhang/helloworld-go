package main

import "fmt"

func main() {
	//A value x of non-interface type X and a value t of interface type T are comparable when values of type X are comparable and X implements T. They are equal if t's dynamic type is identical to X and t's dynamic value is equal to x. ——以上内容来自 https://go.dev/ref/spec#Operators
	//个人理解：
	//1. interface类型变量与非interface类型变量判等时，首先要求非interface类型实现了接口，否则编译不通过（本题接口方法集为空，我们认为所有类型都实现了该接口）
	//2. 满足上一条的前提下，interface类型变量的动态类型、值均与非interface类型变量相同时，两个变量判等结果为true，结合array判等规则，答案为true
	var p [100]int
	var m interface{} = [...]int{99: 0}
	fmt.Println(p == m)
}
