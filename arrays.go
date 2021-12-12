package main

import "fmt"

var array1 [3]int
var array2 = [5]int{1, 2, 3, 4, 5}

func main() {
	array3 := make([]int, 3)

	fmt.Println(array1, array2, array3)
	fmt.Printf("Array 1 Length : %d\n", len(array1))
	fmt.Printf("Array 2 Length : %d\n", len(array2))
	fmt.Printf("Array 3 Length : %d\n", len(array3))
}
