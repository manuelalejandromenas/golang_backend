package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("the value:", m["Answer"])
	
	m["Answer"] = 48
	fmt.Println("the value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("the value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("the value:", v, "Present?", ok)
}
