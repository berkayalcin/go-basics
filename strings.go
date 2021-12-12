package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "lorem ipsum dolor sit amet"
	str2 := str[:5]
	str3 := str[len(str)-4:]

	fmt.Println(str, str2, str3)

	if strings.EqualFold(str2, "LOrEM") {
		fmt.Println("str2 Equal to LOrEM")
	}
}
