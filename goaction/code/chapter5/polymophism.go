package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("sending user email to %v<%v>\n", u.name, u.email)
}

type admin struct {
	name  string
	email string
}

func (a *admin) notify() {
	fmt.Printf("sending admin email to %v<%v>\n", a.name, a.email)
}

func sendNotification(n notifier) {
	n.notify()
}

func main() {
	//create a user value and pas it to send notification
	user1 := user{"User1", "user1@gmail.com"}

	admin1 := admin{"admin1", "admin1@gmail.com"}

	sendNotification(&user1)
	sendNotification(&admin1)
}
