// Flight simulator implementation using Golang Mutex APIs
// The goal is to understand Mutex API and how to use them with
// goroutines.
package main

import (
	"fmt"
	"sync"
)

const noOfFlights = 10

var (
	// wait group to control the program termination.
	wg sync.WaitGroup

	// mutex which simulates the runway
	mutex sync.Mutex
)

func init() {
	// set the Waitgroup with # of goroutines that we are going to create
	wg.Add(noOfFlights)
}

func main() {
	fmt.Println("Flight simulator implementation using Golang Mutex APIs")
	fmt.Printf("Total # of Flight simulated is %d\n", noOfFlights)

	// create flights
	for count := 1; count <= noOfFlights; count++ {
		// run the flights as goroutines
		go flight(count)
	}

	// await for all the flights to land
	wg.Wait()
	fmt.Println("All flights landed successfully....")
}

func flight(number int) {
	// notify the wait group
	defer wg.Done()

	fmt.Printf("Flight %d awaiting for landing\n", number)
	mutex.Lock()
	{
		fmt.Printf("Flight %d received permission for landing\n", number)
		fmt.Printf("Flight %d has been landed safely and successfully\n", number)
	}
	mutex.Unlock()
}
