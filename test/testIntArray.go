package main

import "fmt"

func main() {
	//  a的ASCII码数值是97，b的ASCII码数值是98，c的ASCII码数值是99
	m := [...]int{
		'a': 1,
		'b': 2,
		'c': 3,
	}
	m['a'] = 3
	fmt.Println(len(m))
}
