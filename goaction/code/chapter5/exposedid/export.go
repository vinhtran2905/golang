package main

import (
	"fmt"

	"github.com/vinhtran2905/golang/goaction/code/chapter5/exposedid/entities"
)

func main() {
	fmt.Println("Test the export identifer User from entities package")

	//This will cause an error 'implicit assignment of unexported field 'name' in entities.User literal'
	//because the name and email are unexported fields
	// u := entities.User{"Vinh", "v@gmail.com"}
	// fmt.Printf("sending email to user %v<%v>\n", u.name, u.email)

	a := entities.Admin{
		Wage: 100.5,
	}

	//Set the exported fileds from the unexported inner type
	a.Name = "admin"
	a.Email = "admin@gmail.com"
	fmt.Printf("sending email to user %v<%v>\n", a.Name, a.Email)

}
