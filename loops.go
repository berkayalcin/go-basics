package main

import (
	"fmt"
	"time"
)

func main() {
	slice2 := make([]int, 0, 3)
	slice2 = append(slice2, 1, 2, 3, 4, 5, 6, 7)

	for i, value := range slice2 {
		fmt.Printf("Index : %d Value : %d \n", i, value)
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("I : %d\t", i)
	}

	for i, value := range array2 {
		fmt.Printf("Index : %d Value : %d \n", i, array2[i])
		value += 1
		fmt.Printf("Index : %d Value : %d \n", i, array2[i])
	}

	// While Loop
	c := time.After(5 * time.Second)
	for {
		b := false
		select {
		case <-c:
			b = true
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}

		if b {
			break
		}
	}

}
