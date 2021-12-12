package main

import (
	"errors"
	"fmt"
	"log"
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

func (u User) followerCount() int {
	return len(u.Followers)
}

func (u *User) addFollower(user User) error {
	if u.Followers == nil {
		u.Followers = []User{}
	}

	if u.followerCount() == 10 {
		return errors.New("max User Length Exceeded")
	}

	u.Followers = append(u.Followers, user)
	return nil
}

func main() {
	user := User{Name: "Berkay", Surname: "Yalçın", Followers: []User{}}
	user.addFollower(User{Name: "Unknown", Surname: "User"})
	for i := 0; i < 10; i++ {
		if err := user.addFollower(User{Name: "Unknown", Surname: "User"}); err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Printf("Follower Count %d\n", user.followerCount())
}
