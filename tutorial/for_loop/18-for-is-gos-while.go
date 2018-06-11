package main

import "fmt"
// While is spelled for in Go

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
