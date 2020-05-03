package main

import (
	"fmt"

	"github.com/vinhtran2905/golang/goaction/code/chapter5/counters/"
)

func main() {
	counter := counters.AlertCounter(10)
	fmt.Printf("Counter: %v\n", counter)
}
