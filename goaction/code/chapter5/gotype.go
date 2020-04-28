package main

import (
	"fmt"
	"reflect"
)

//2 ways to defined new type in GO: struct or takine an existing type and use it as the type specification for the new type
type user struct {
	name    string
	email   string
	age     int
	married bool
}

func main() {

	fmt.Printf("User defined types \n")

	var u user
	fmt.Printf("Type user %v \n", u)

	u2 := user{
		name:    "Vinh",
		email:   "vinh@gmail.com",
		age:     13,
		married: false,
	}

	fmt.Printf("u2 %v\n", u2)
	fmt.Printf("Name of u2 is %v\n", u2.name)
	fmt.Printf("Is he married? %v\n", u2.married)

	u3 := user{"Minh", "minh@gmail.com", 10, true}
	fmt.Printf("u3 %v\n", u3)
	fmt.Printf("Name of u2 is %v\n", u3.name)
	fmt.Printf("Is he married? %v\n", u3.married)

	//Define a struct field based on other struct type
	type admin struct {
		u      user
		salary float64
	}

	a1 := admin{u2, 100.9}
	fmt.Printf("admin a1 is %v\n", a1.u)
	fmt.Printf("admin a1 has salary is %v\n", a1.salary)

	a2 := admin{
		salary: 200,
		u: user{
			name:    "Lisa",
			married: true,
		},
	}
	fmt.Printf("admin 2 has user %v\n", a2.u)
	fmt.Printf("admin a2 has salary is %v\n", a2.salary)

	type newint int64
	//Go does not consider newint and int64 are the same type. They are two distinct and different types
	var count newint
	count = 1
	fmt.Printf("count has type %v and value %v\n", reflect.TypeOf(count), count)
	fmt.Printf("count has type %T\n", count)

	var i int64
	//count = i this will cause error since i and count are different types
	//You must cast type of i to type of count before assigning i to count
	count = newint(i)
	fmt.Printf("count type is %T has value is %v\n", count, count)
	fmt.Printf("i type is %T has value is %v\n", i, i)

	//Methods provide a way to add behavior to user-defined types.

	u3.notify()
	u3.changeEmail("newminh@gmail.com")
	u3.notify()
	u3.changeName("NewNAME")
	u3.notify()

	lisa := &user{"Lisa", "lisa@gail.com", 56, true}
	//we can call methods that are declared with a value receiver using a pointer
	lisa.notify()
	//(*lisa) is dereferenced
	(*lisa).notify()

}

//parameter betwen func and funcion name is called a RECEIVER and binds the function to the specified type.
//When a function has a receiver, that function is called a method
func (u user) notify() {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)

}

//This method is using a VALUE RECEIVER.
func (u user) changeEmail(email string) {
	u.email = email
}

//This method is using a POINTER RECEIVER
func (u *user) changeName(name string) {
	u.name = name
}
