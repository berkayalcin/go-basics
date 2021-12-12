package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type User interface {
	followerCount() int
	addFollower(user *User) error
}

type user struct {
	Name      string `json:"name"` // tags
	Surname   string `json:"-"`
	Followers []user `json:"followers,omitempty"`
	Likes     []struct {
		date time.Time
	} `json:"likes,omitempty"`
}

func (u user) followerCount() int {
	return len(u.Followers)
}

func (u *user) addFollower(follower user) error {
	if u.Followers == nil {
		u.Followers = []user{}
	}

	if u.followerCount() == 10 {
		return errors.New("max User Length Exceeded")
	}

	u.Followers = append(u.Followers, follower)
	return nil
}

func greaterThan(u1, u2 user) bool {
	return u1.followerCount() > u2.followerCount()
}

var u User

func main() {
	u := user{Name: "Berkay", Surname: "Yalçın", Followers: []user{}}
	u.addFollower(user{Name: "Unknown", Surname: "User"})
	for i := 0; i < 9; i++ {
		if err := u.addFollower(user{Name: "Unknown", Surname: "User"}); err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Printf("Follower Count %d\n", u.followerCount())
	fmt.Printf("Is User 1 Follower Count Greater than User 2 %s", greaterThan(u, u))
}
