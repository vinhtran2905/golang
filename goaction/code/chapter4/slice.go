package main

import "fmt"

func main() {
	//declare a slice of string has length of 3 by using make
	s1 := make([]string, 3)

	//declare a slice of string has length of 3 and capacity of 5
	s2 := make([]string, 3, 5)

	//
	s1 = append(s1, "hello")
	s2 = s1
	fmt.Println(s2)

	s1 = append(s1, "world")
	fmt.Println(s1)
	fmt.Println(s2)

	//declare slice with a slice literal
	s3 := []string{"test1", "test2"}
	fmt.Println(s3)

	//declare a slice with index position so that this slice has 100 elements
	s4 := []string{99: "test100"}
	fmt.Println(s4)

	// declaration differences between arrays and slices
	array := [3]int{1, 2, 3}
	slice := []int{1, 2, 3}
	fmt.Printf("array type: %v \n", array)
	fmt.Printf("slice type: %v", slice)

	//Nil and empty slice
	//nil slice is usedfull when you want to represent a slice that does not exist such as exception
	var nilslice []int //nil pointer, length of zero and capacity of zero
	//nilslice := make([]int)
	// nilslice := []int
	fmt.Printf("\nNil slice %v", nilslice)

	//empty slice is useful when you want to represent an empty collection such as when database query returns a zero results
	//empty pointer, length of zero and capacity of zero
	emptyslice := make([]int, 0)
	//  emptyslice := []int{0}
	fmt.Printf("\nempty slice %v\n", emptyslice)

	//assign and slicing slice
	s5 := []int{10, 20, 30, 40, 50}
	s5[1] = 21
	fmt.Println(s5)
	//s6 will have length 2 and capacity 4
	//because: 3-1 = 2
	//underlying array s5 has 5 capacity, therefore the capacity of the new slicing array s6 has 5-1 = 4 capacity
	//formula: slice[i:j] with an underlying array of capacity k
	//length = j - i
	//capacity = k - i

	s6 := s5[1:3]
	fmt.Printf("Slicing: %v\n", s6)
	fmt.Printf("Capacity of slicing slice %v\n", cap(s6))

	//changes made to the shared section of the underlying array by one slice can be seen by the other slice

	s6[1] = 100
	fmt.Printf("original slice %v\n", s5)
	fmt.Printf("new slice slice %v\n", s6)
	//Having capacity is great, but if you cannt incorporate it into your slice's length. Luckily, GO makes this easy when you use the built-in function append

	//Growing slice: once of the advantages of using a slice over using an array is that you can grow the capacity of your slice as needed. Go takes care of all the operational details when you use the builtin function append
	s6 = append(s5, 60)
	fmt.Printf("After appending 2 elements into the new slice \n")
	fmt.Printf("original slice %v\n", s5)
	fmt.Printf("new slice slice %v\n", s6)

	s7 := []int{10, 20, 30, 40, 50}
	s8 := s7[1:3]
	s8 = append(s8, 60)
	fmt.Printf("After appending 2 elements into the new slice. original and new slice will get the changes since the underlying array has changed \n")
	fmt.Printf("original slice %v\n", s7)
	fmt.Printf("new slice slice %v\n", s8)

	//expand capacity of a slice. if nb of enderlying array is less than 1000, capacity will extend unto double of the current capacity
	//if nb of elements of the underlying array is more than 1000, the capacity will be increased to 1.25 or 25%
	s9 := []int{1, 2, 3, 4, 5}
	s10 := append(s9, 6)
	fmt.Printf("len and capacity of original slice: %v %v \n", len(s9), cap(s9))
	fmt.Printf("len and capacity of new slice slice: %v %v \n", len(s10), cap(s10))

	//slice with the third index
	s11 := []string{"apple", "orange", "plum", "banana", "grape"}
	s12 := s11[2:3:4]
	//newslice[i:j:k]
	//len = j-i = 3-2 =1,
	//cap= k-i = 4-2 =2 < cap of s11= 5
	fmt.Printf("len and cap of s12: %v, %v\n", len(s12), cap(s12))

	//error when setting capacity is larger than existing capacity. This reason we should use append to expand the slice capacity

	// s13 := s11[2:3:6]
	// fmt.Printf("len and cap of s12: %v, %v\n", len(s13), cap(s13))

	//Detach the new slice from its original source array makes it safe to change
	//without the third index, appending KIWI to our slice would've changed the value of Banana in index 3 of the underlying array
	source := []string{"apple", "orange", "plum", "banana", "grape"}
	destination := source[2:3:3]
	fmt.Printf("destination slice %v with len %v and cap %v\n", destination, len(destination), cap(destination))

	destination = append(destination, "KIWI")
	fmt.Printf("source slice %v with len %v and cap %v\n", source, len(source), cap(source))
	fmt.Printf("destination slice %v with len %v and cap %v\n", destination, len(destination), cap(destination))

	//variadic function by using ... operator to append all the elements of one slice to another
	s14 := []int{2, 3}
	s15 := []int{4, 5}
	s15 = append(s14, s15...)
	fmt.Printf("s15: %v\n", s14)
	fmt.Printf("s15: %v\n", s15)

	//iterating over slices
	for index, value := range s15 {
		fmt.Printf("Index %v Value %v\n", index, value)
	}

	//if you dont need the index, use can ignore it by using underscore

	for _, value := range s15 {
		fmt.Printf("Value %v\n", value)
	}

	//multi dimensinoal slices
	//create a slice of slice of integers

	ms := [][]int{{10}, {100, 200}}
	fmt.Printf("multi dimentional slice: %v\n", ms)
	ms[0] = append(ms[0], 20)
	fmt.Printf("slice[0]: %v\n", ms[0])
	fmt.Printf("multi dimentional slice: %v\n", ms)

	//passing the slices between functions requires nothing more than passing the slice by value. Slice is requires 24 bytes: 8 bytes for pointer, 8 bytes for lenghth, and 8 bytes for capacity.
	//this is beautiful of slice. You dont need to pass pointers around and deal with complicated.
	//You just create copies of your slices make the changes you need and then pass a new copy back
	s16 := make([]int, 1e6)
	foo(s16) // just pass 24 bytes into the foo instead of pass 1e6 elements into the function
}

func foo(slice []int) []int {
	return slice
}
