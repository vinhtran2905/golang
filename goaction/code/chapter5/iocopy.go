package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

//use interfaces to concate and then stream data to standard out
func main() {
	var b bytes.Buffer
	//write a string to the buffer
	b.Write([]byte("Hello"))

	//Fprintf to concatenate a string to the buffer
	fmt.Fprintf(&b, "World")

	io.Copy(os.Stdout, &b)

}
