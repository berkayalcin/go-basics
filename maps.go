package main

import "fmt"

var map1 map[int]string

func main() {
	map1 = make(map[int]string)
	map2 := make(map[int]int)
	map3 := make(map[int]int, 3)
	map1[0] = "1"
	map1[1] = "A"

	map2[0] = 1
	map2[1] = 1

	map3[0] = 1
	map3[1] = 1
	map3[2] = 1
	map3[3] = 1

	fmt.Println(map1, map2, map3)
	fmt.Printf("Map 1 Length : %d \n", len(map1))
	fmt.Printf("Map 2 Length : %d \n", len(map2))
	fmt.Printf("Map 3 Length : %d \n", len(map3))
}
