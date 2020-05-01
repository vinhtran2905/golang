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
	u3.changeEmail("newminh@gmail.com") // This will not change new email because value receiver operated against a copy of the value
	u3.notify()
	u3.changeName("NewNAME") // This will change new NAME since POINTER RECEIVER uses the value used to make a call is shared with the method
	u3.notify()

	//(*lisa) is dereferenced. We can call a method that is declared with a value receiver using a pointer
	lisa := &user{"Lisa", "lisa@gail.com", 56, true}
	lisa.notify()
	(*lisa).notify()
	lisa.changeEmail("new email") // This will not change new email since it operated against a copy of the value
	lisa.notify()

	lisa.changeName("NewLisa") //The name will be changed
	(*lisa).notify()

	//We can also call the pointer receivcer method by using a value
	u3.changeName("NewNameNew") //underneath Go is using (&u3).changeName("NewNameNew")
	u3.notify()
	(&u3).changeName("NewNameNewNew")
	u3.notify()

	//Value receiver operates on a copy of the value used to make the method call while poiter receiver operates on the actual value
	//Guidelines whether to use a value or a pointer receiver:
	//After declaring a new type, try to answer this question before declaring methods for the type.
	//Does adding or removing something from a value of this type need to create a new value or mutate the existing one?
	//If the answer is create a new value, then use value receiver. If the answer is mutate the value then use the pointer receiver.
	//This also applies to how values of this type should be passed to other parts of your program.
	//This idea is to not focus on what the method is doing with the value, but to focus on what the nature of the value is

	//1. BUILT-IN TYPE: Numeric, string and boolean. Passing values of these ttypes to functions and methods, a copy of the value should be passed

	//2. REFERENCE TYPES: slice, map, channel, interface and function types
	//When you declared a variable from one of these types, the value that's created is called a header value.
	//Each reference type contains a POINTER to an underlying data structure and a set of unique fields that are used to manage the underlying data structure
	//You never share reference type values because the header value is designed to be copied. The header value contains a POINTER, therefore
	//you can pass a copy of any reference type value and share the underlying data structure.REference type values are treated like primitive data values
	//E.g: type IP []byte
	// func (ip IP) MarshallText() ([]byte, error) {
	// 	if len(ip) == 0 || ip == nil {
	// 		return []byte(""), nil
	// 	}
	// if len(ip) != IPv4len && len(ip) != IPv6len {
	// 	return nil, errors.New("invalid IP")
	// }
	// return []byte(ip.String()), nil
	// }

	// func ipEmptyString(ip IP) string {
	// 	if len(ip) == 0 {
	// 		return ""
	// 	}
	// 	return ip.String()
	// }

	//3 STRUCT TYPES: can represent data values that could have either a primitive or non primitive nature.
	//In most cases, adding or removing something freom the value of the STRUCT type should mutate the value.
	//When this is the case, you want to use a pointer to share the value with the rest of the program that needs it.

	// type File struct {
	// 	*flie
	// }

	// type file struct {
	// 	fd int
	// 	name string
	// 	dirInfo *dirInfo
	// 	nepipe int32
	// }

	//The decision to use a VALUE or POINTER RECEIVER should not be based on whether the method is mutating the receiving value.
	//The decision should be based on the nature of the type.
	//One exeception to this guideline is when you need the flexibility that value type receivers provide when working with INTERFACE values.
	//In this case you may choose to use a VALUE receiver even though the nature of the type is NONPRIMITIVE

	//4 INTERFACE TYPE
	//Polymorphism is the ability to write code that can take on differnt behavior through the implementation of types.
	//Onces a type implements an interface, an entire world of functionality can  be opened up to values of that type
}

//parameter between func and funcion name is called a RECEIVER and binds the function to the specified type.
//When a function has a receiver, that function is called a method
func (u user) notify() {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)

}

//This method is using a VALUE RECEIVER. When you declare a method using a value receiver,
//the method will always be operating against a copy of the value used to make the method call.
func (u user) changeEmail(email string) {
	u.email = email
}

//This method is using a POINTER RECEIVER
func (u *user) changeName(name string) {
	u.name = name
}
