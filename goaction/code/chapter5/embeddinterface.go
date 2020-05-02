package main

import "fmt"

type notifer interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("sending email to user %v<%v>\n", u.name, u.email)
}

type admin struct {
	user
	role string
}

func sendNotifier(n notifer) {
	n.notify()
}

func main() {
	lisa := user{"Lisa", "l@g.com"}
	ad := admin{
		user: lisa,
		role: "admin",
	}

	ad.user.notify()
	fmt.Println("admin call notify directly")
	ad.notify()

	//sendinf the admin user a notification.
	//The embedded inner type's implementation of the interface is promoted to the outer type
	//Thanks to inter type promotion, the implementation of the interface by the inner type has been
	//promoted up the the outer type. That means the outer type now implements the interface, thanks to
	// the inner type
	sendNotifier(&ad)
	sendNotifier(&(ad.user))

	//if the outer type implements the interface method, what will be happened?
	//The inner's method is not promoted since the outer already have its own method
}
