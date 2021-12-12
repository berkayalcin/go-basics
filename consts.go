package main

import "fmt"

// Iota - Enumarable Number
const (
	const1 = iota
	const2 = iota
	const3 = iota
)
const const4 = "Constant"

func main() {
	fmt.Println(const1, const2, const3, const4)
}
