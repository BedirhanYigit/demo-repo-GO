package main

import (
	"fmt"
	"sync"
)

func squareWorker(id int, input <-chan int, output chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range input {
		result := num * num
		output <- result
		fmt.Printf("Worker %d: %d * %d = %d\n", id, num, num, result)
	}
}

func main() {
	/*	var name = "Go conference"
		const conferenceTicket = 50
		fmt.Println("Welcome to our", name, "booking app")
		fmt.Println("Get your tickets here to attend")
		fmt.Println("Tickets cost", conferenceTicket, "Euro")

		// now lets add a new line to check from github
		fmt.Println("This is the new line that has been recently added")
	*/

	/*
		// printing a string with syntactic sugar :=
		name := "Syntactic Sugar Bedirhan"
		fmt.Println(name)

		// printing a string with a normal variable
		var name1 string = "Normal Bedirhan"
		fmt.Println(name1)

		// using the var feature and changing the strings
		name = "Syntactic Sugar Bedirhan has been changed"
		name1 = "Normal Bedirhan has been changed"
		fmt.Println(name)
		fmt.Println(name1)

		// practise with for and while loops, print numbers from 1 to 10

		// classic for loop
		for i := 1; i < 11; i++ {
			fmt.Println(i)
		}
		fmt.Println("Classic for loop has been finished")

		// for loop with a parameter that has been already created
		j := 1
		for j < 11 {
			fmt.Println(j)
			j++
		}
		fmt.Println("for loop with a schon created variable has been finished")

		// while loop
		i := 1
		for {
			fmt.Println(i)
			i++
			if i == 11 {
				break
			}

		}
		fmt.Println("While loop with a for loop has been finished")

		// pointers
		var someNumber int = 30
		var somePointer *int = &someNumber

		// printing out the value that somePinter shows
		fmt.Println("somePointer points the value of someNumber:", *somePointer)
		fmt.Println("the value of someNumber:", someNumber)

		//printing out the storage of someNumber and then somePointer to point out that they have different storage in ram
		fmt.Println("Address of someNumber:", &someNumber)
		fmt.Println("Address of somePointer:", &somePointer)

		// arrays with fixed size, and printing the values with
		var intArray [4]int
		intArray[0] = 20
		intArray[1] = 30
		intArray[2] = 40
		intArray[3] = 50
		for i := 0; i < len(intArray); i++ {
			fmt.Printf("Element %d: %d\n", i, intArray[i])
		}

		// slice (dynamic sized array)
		intArray1 := []int{1, 2, 3, 4, 5}
		for i, value := range intArray1 {
			fmt.Printf("Element %d: %d\n", i, value)
		}
	*/

	// Go routines example

	const numWorkers = 5
	const numTasks = 10

	// Create channels for communication between main and worker goroutines
	inputChannel := make(chan int, numTasks)
	outputChannel := make(chan int, numTasks)

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Launch worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go squareWorker(i, inputChannel, outputChannel, &wg)
	}

	// Send tasks to the workers
	for i := 1; i <= numTasks; i++ {
		inputChannel <- i
	}

	// Close the input channel to signal that no more tasks will be sent
	close(inputChannel)

	// Wait for all workers to finish
	wg.Wait()

	// Close the output channel to signal that no more results will be received
	close(outputChannel)

	// Collect and print the results from the output channel
	for result := range outputChannel {
		fmt.Printf("Main: Received result %d\n", result)
	}

}
