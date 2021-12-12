package main

import (
	"fmt"
	"strings"
)

type String *string

func main() {
	var rstr String
	var pstr string

	// Nil - Reference Type
	fmt.Println(rstr)
	fmt.Println(pstr)

	pstr = "Go Türkiye"
	rstr = &pstr // Pointer Address

	fmt.Println(rstr)
	fmt.Println(pstr)

	pstr = "Go Türkiye Eğitim Kampı"
	fmt.Println(pstr)
	fmt.Println(*rstr)

	*rstr = "go Türkiye Eğitimi"
	fmt.Println(pstr)
	fmt.Println(*rstr)

	replace(rstr)
	fmt.Println(pstr)
	fmt.Println(*rstr)
}

func replace(s String) {
	*s = strings.Replace(*s, "go", "GO", 1)
}
