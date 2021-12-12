package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Name      string `json:"name"` // tags
	Surname   string `json:"-"`
	Followers []User `json:"followers,omitempty"`
	Likes     []struct {
		date time.Time
	} `json:"likes,omitempty"`
}

type Like struct {
	Date time.Time
}

func main() {
	user := User{Name: "Berkay", Surname: "Yalçın", Followers: []User{
		{Name: "Unknown", Surname: "User"},
	}}

	arr, err := json.Marshal(user)
	fmt.Printf("Error : %s\n", err)
	fmt.Println(string(arr))
	fmt.Println(user)
}
