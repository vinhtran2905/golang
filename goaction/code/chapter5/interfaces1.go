package main

import (
	"fmt"
)

//interface is type that use declared behavior.
//This behavior is never implemented by the interface type directly but instead by user-defined types via methods.

type notifier interface {
	notify()
}

//user define a user in the program
type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("sending user email to %s<%s>", u.name, u.email)
}

func main() {
	lisa := user{"Lisa", "lisa@gmail.com"}
	sendNotification(&lisa)
}

func sendNotification(n notifier) {
	n.notify()
}
