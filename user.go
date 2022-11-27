package main

import (
	"fmt"
	"time"
)

type User struct {
	id string
}

func (u *User) update(itemName string) {
	time.Sleep(time.Second * 2)
	fmt.Printf("sending email to user with id %s, for item %s is now available\n", u.id, itemName)
}

func (u *User) getID() string {
	return u.id
}
