package main

import (
	"fmt"
	"time"
)

func main() {
	s, err := timer(time.After(5*time.Second), "Go", "Türkiye", "Eğitimi")
	fmt.Println(s, err)
}

func timer(c <-chan time.Time, message ...string) (s string, err error) {
	defer fmt.Println("Going to next step")
	defer fmt.Println("Timer Is Done")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	s = "Go Türkiye"
	err = nil
	for {
		select {
		case <-c:
			fmt.Println(message)
			return
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
			if time.Now().Day() == 12 {
				panic("Unexpected error occured")
			}
		}
	}
}
