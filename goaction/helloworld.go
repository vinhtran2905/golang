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
		str = "golang world"
	}
	return helloPrefix + str
}
