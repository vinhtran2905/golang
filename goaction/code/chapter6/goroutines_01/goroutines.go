package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//allocate 1 logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	//wg is used to wait for the program to finish
	//add a count of two, one for each gorouting
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start Goroutines")

	//Declare an anonymous function and create a goroutine
	go func() {
		//Scheduler the call to Done to tell main we are done
		defer wg.Done()
		for count := 0; count < 1; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	//wait for the goroutines to finish
	fmt.Println("waiting to finish")
	wg.Wait()
	fmt.Println("\n Terministing program")

	// The amount of time it takes the first goroutin to finish displaying the alphabet is so small
	//that it can complete its work before the scheduler swaps it out for the decond goroutine.
	//This is why you see the entire alphabet in capital letters first and the nin lowercase letters second
	//The 2 goroutines we created ran concurrently one after the other, performing their individual task of displaying the alphabet
	//Once the 2 anonymous functinos are created as goroutines, the code in main keeps running.
	//This means that the main function can return before the goroutines complete their work.
	//If this happens the program will terminate before the goroutines have a chance to run.
	//Therefore the main function waits for both goroutines to complete their work by using WaitGroup
}
