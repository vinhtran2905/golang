//The ability to apply visibility ruls to the identifiers you declare is critical for good API design.
//Go suports the exporting of identifiers from a package to provide this functionality
package main

func main() {
	counter := counters.alertCounter(10)
	fmt.Printf("Counter:: %d\n",counter)
}