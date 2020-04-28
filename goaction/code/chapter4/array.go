package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//declare and initialize an array by using array literal

	arr2 := [5]int{10, 20, 30, 40, 50}

	for j := 0; j < 5; j++ {
		fmt.Println(arr2[j])
	}

	//length of an array based on the nb of the elements that are initialized
	arr3 := [...]int{100, 200, 300, 400, 500}
	for k := 0; k < len(arr3); k++ {
		fmt.Println(arr3[k])
	}

	//array with initial elements such as initialize first element = 10 and the fouth element = 40
	arr4 := [5]int{3: 40, 0: 10}
	for l := 0; l < len(arr4); l++ {
		fmt.Print(arr4[l], " ")

	}
	//array pointer element
	arr5 := [5]*int{0: new(int), 1: new(int)}
	*arr5[0] = 5
	*arr5[1] = 10
	fmt.Println("\n**********array pointer")
	for m := 0; m < len(arr5); m++ {
		fmt.Println(arr5[m])
	}
	fmt.Printf("data at %v is value %v", arr5[0], *arr5[0])

	//assign an array to an array
	fmt.Println("\n**********Array assigned an array")
	var arr6 [5]string
	arr7 := [5]string{"s0", "s1", "s2", "s3", "s4"}
	arr6 = arr7
	for n := 0; n < len(arr6); n++ {
		fmt.Print(arr6[n], " ")
	}
	arr7[4] = "new string"
	fmt.Println("\n")
	for n := 0; n < len(arr7); n++ {
		fmt.Print(arr7[n], " ")
	}

	//assign one array of pointers to another
	var pa1 [3]*string

	pa2 := [3]*string{new(string), new(string), new(string)}

	*pa2[0] = "st1"
	pa1 = pa2
	fmt.Println("\n", *pa1[0])

	*pa2[0] = "newst"
	fmt.Println(*pa1[0])
	fmt.Println(*pa2[0])

	//2 dimension array
	var arr2d1 [4][2]int
	arr2d2 := [4][2]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}}
	fmt.Println(arr2d2)
	arr2d1 = arr2d2
	fmt.Println(arr2d1)
	arr2d2[0][0] = 10
	arr2d2[0][1] = 20
	fmt.Println(arr2d2)

	//array 2 d with pointer

	arr2d3 := [3][2]*string{{new(string), new(string)}, {new(string), new(string)}, {new(string), new(string)}}
	fmt.Print(arr2d3)
}
