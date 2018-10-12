package main

import (
	"fmt"
)

const helloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("world"))

}

func Hello(str string) string {
	if str == "" {
		str = "world"
	}
	return helloPrefix + str
}
