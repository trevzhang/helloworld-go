package main

import (
	"encoding/json"
	"fmt"
)

// Define a struct type for the JSON structure
type Result struct {
	Age int `json:"age"`
}

type ResultPointer struct {
	Age *int `json:"age"`
}

func main() {
	// 1. Age int = null
	jsonData := `{"age": null}`
	res := &Result{}
	if err := json.Unmarshal([]byte(jsonData), res); err != nil {
		panic(err)
	}
	age := res.Age
	fmt.Println("Age:", age, "\n") // Output will be Age: 0 since the default for int is 0

	// 2. Age int = 12
	jsonData2 := `{"age": 12}`
	res2 := &Result{}
	if err := json.Unmarshal([]byte(jsonData2), res2); err != nil {
		panic(err)
	}
	age2 := res2.Age
	fmt.Println("Age:", age2, "\n") // Output will be Age: 12

	// 3. Age *int = null
	jsonData3 := `{"age": null}`
	resPointer := &ResultPointer{}
	if err := json.Unmarshal([]byte(jsonData3), resPointer); err != nil {
		panic(err)
	}
	age3 := resPointer.Age
	fmt.Println("Age:", age3, "\n") // Output will be Age: <nil>

	// 4. Age *int = 12
	jsonData4 := `{"age": 12}`
	resPointer2 := &ResultPointer{}
	if err := json.Unmarshal([]byte(jsonData4), resPointer2); err != nil {
		panic(err)
	}
	age4 := resPointer2.Age
	fmt.Println("*Age:", age4, "\n") // Output will be Age: <nil>
	fmt.Println("Age:", *age4, "\n") // Output will be Age: <nil>
}
