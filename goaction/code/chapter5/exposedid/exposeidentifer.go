package main

import (
	"fmt"

	//manuall typing
	"github.com/vinhtran2905/golang/goaction/code/chapter5/exposedid/vinh"
)

func main() {
	counter := vinh.New(4)
	fmt.Printf("counter: %v\n", counter)

	t := vinh.Timer(5.6)
	fmt.Printf("timer %v\n", t)
}
