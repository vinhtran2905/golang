//Go allows you to take existing types and both extend and change thier behavior.
//This capacity is important for code reuse and for changing the behavior of an existing type to suit a new need
//This is accomplished through type embedding. It works by taking an existing type and declaring the type within
//the declaration of a new struct type. The type that is embedded is then called an innner type of the new outer type

//This is different with the composition

package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
}

//user type is an inner type of the outer type admin
type admin struct {
	user //Embedded type
	role string
}

func main() {

	admin1 := admin{
		user: user{"Vinh", "vinh@gmail.com"},
		role: "director",
	}

	admin1.user.notify()
	//why user does not using a pointer receicer?
	//testing
	user2 := &user{"user2", "u2@gmail.com"}
	user2.notify()

	user3 := user{"user3", "u3@gmail.com"}
	user3.notify()

	//thanks to inner type promotion, the notify method can also be accessed directly from the ad variable
	admin1.notify()

}
