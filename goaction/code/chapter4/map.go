package main

import "fmt"

//Map is a data structure that provides you with an unordered collection of key/value pairs.
//The strength of a map is its ability to retrieve data quickly based on the key.
//A key work like an index pointing to the value you associate with that key
//Every iteration over a map could return a different order. This is because a map is implemented using a hash table.
//Map's hash table contains a collection of buckets. Keys are converted into a numeric value by a hash function.
// This numeric value is then used to select a bucket or storing or finding the specific key/value pair.
//The numeric has 2 parts: Low Order Bits (LOB) is used to select the bucket.
//There are two data structures that contain the data for the map.
//First there's an array with the top eight High Order Bits (HOB) from the same hash key that was used to select the bucket
//This array distinguishes each individual key/value pair stored in the respective bucket.
//Second, there's an array of bytes that stores the key/value pairs.
//The byte array packs all the keys and then all the values together ror the respective bucket.
// The packing of the key/value pairs is implemented to minimize the memory required for each bucket

func main() {
	fmt.Printf("This program guide how to use map in Go \n")

	//Creating and initializing a map
	//1. Declare a map by make

	m1 := make(map[string]int)
	//s := make([]int,5) //to define a slice by make

	fmt.Printf("map 1: %v\n", m1)

	//2. Initialize a map
	m2 := map[string]float64{"AMZ": 2333.9, "DIS": 109.8}
	fmt.Printf("m2: %v\n", m2)

	//declare map by using a map literal
	//map that contains a slice of strings.
	m3 := map[int][]string{1: {"red", "black"}, 2: {"apple", "banana", "plum"}}
	fmt.Printf("m3: %v \n", m3)

	//error when key is a collections
	// m4 := map[[]string]int{}
	// fmt.Printf("Error when define a map contains key is a collection %v\n", m4)

	//nil vs empty map
	//nil map
	var m5 map[int]string
	// m5[1] = "hello" // cause a panic exception since a nil map cannot be used to store key/value pairs
	fmt.Printf("a nil map: %v\n", m5)
	//empty map
	m6 := map[int]string{}
	fmt.Printf("a empty map: %v \n", m6)
	m6[1] = "hello"
	m5 = m6
	fmt.Printf("a nil map after assign to a map: %v\n", m5)

	//Retrieve the value for a key "1"

	value, exists := m5[1]
	if exists {
		fmt.Printf("Value of key 1 of map m5 is:  %v\n", value)
	} else {
		fmt.Printf("There is no exist m5[2]\n")
	}

	//dont care about flag exist we can use the underscore
	value2, _ := m5[2]
	fmt.Printf("value is %v\n", value2)

	//Iterating over a map using for range
	for k, v := range m3 {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}

	//Removing an item from a map
	delete(m3, 1)
	fmt.Printf("Map after removed %v\n", m3)

	//Passing a map between function does not make a copy of the map. In fact, you can pass a map to a function and
	//make changes to the map and the changes will be reflected by all references to the map

	removeItem(m3, 2)
	fmt.Printf("m3 after adding an item %v\n", m3)
}

func removeItem(m map[int][]string, k int) {
	delete(m, k)
}
