package main

import (
	"fmt"

	//manually typing
	"github.com/vinhtran2905/golang/goaction/code/chapter5/exposedid/vinh"
)

func main() {
	counter := vinh.New(4)
	fmt.Printf("counter: %v\n", counter)

	// t := vinh.Newtimer(5.4)
	// fmt.Printf("timer %v\n", t)

	// t := vinh.Newtimer(10)
	// fmt.Printf("counter: %v\n", t)

}
