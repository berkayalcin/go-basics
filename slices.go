package main

import "fmt"

var slice1 []int

func main() {
	slice2 := make([]int, 0, 3)
	slice2 = append(slice2, 1, 2, 3, 4, 5, 6, 7)

	fmt.Println(slice1, slice2)
	fmt.Printf("Slice 1 Length : %d Capacity : %d \n", len(slice1), cap(slice1))
	fmt.Printf("Slice 2 Length : %d Capacity : %d \n", len(slice2), cap(slice2))
}
