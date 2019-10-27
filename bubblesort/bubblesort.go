package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// The goal of this assignment is to create a routine that sorts a series of ten integers using the bubble sort algorithm.
// Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers.
// The program should print the integers out on one line, in sorted order, from least to greatest.
// Use your favorite search tool to find a description of how the bubble sort algorithm works.
// As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
// The BubbleSort() function should modify the slice so that the elements are in sorted order.
// A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice.
// You should write a Swap() function which performs this operation.
// Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice.
// The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
func main() {
	// Get an initial test slice of 10 integers
	theSliceToBeSorted := getInitialSlice()
	fmt.Printf("So far the initial slice built is %v:\n", theSliceToBeSorted)

	// Sort the slice using the Bubble Sort Algorithm
	BubbleSort(theSliceToBeSorted)

	// Show the resulting sorted slice
	fmt.Println("\n=============================================================================")
	fmt.Printf("The resulting Slice status is %v\n", theSliceToBeSorted)
	fmt.Println("=============================================================================")
	os.Exit(0)
}

// getInitialSlice is an auxiliary function in charge of requesting the user up to 10 integers
// and return an Slice of those integers entered
func getInitialSlice() []int {

	fmt.Println("========================================================")
	fmt.Println(" You are now being asked to introduce up to 10 integers ")
	fmt.Println("========================================================")

	var theSlice = make([]int, 10, 10)
	var cursor int = 0

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Please enter integer #%d of 10 --> ", (cursor + 1))
		in.Scan()
		anInt, err := strconv.Atoi(in.Text())
		if err != nil {
			fmt.Println("Not an integer, please try again.")
		} else {
			theSlice[cursor] = anInt
			if cursor++; cursor >= 10 {
				break
			}
		}
	}
	return theSlice
}

// BubbleSort takes a slice of integers as an argument and sorts its content
// using the Bubble Sort algorithm.
func BubbleSort(aSliceToBeSorted []int) {

	fmt.Println("\n================================")
	fmt.Println("Entering BubbleSort Function")
	fmt.Println("================================")

	fmt.Printf("So far the initial state of the slice received is %v:\n", aSliceToBeSorted)

	var loopCount = 0
	n := len(aSliceToBeSorted)
	infiniteLoopProtection := 100 * n * (n - 1) / 2
	isSorted := false

	for !isSorted {

		numberOfSwaps := 0
		loopCount++

		for i := 0; i < (len(aSliceToBeSorted) - 1); i++ {

			if aSliceToBeSorted[i] > aSliceToBeSorted[i+1] {
				Swap(aSliceToBeSorted, i)
				numberOfSwaps++
			}
		}

		fmt.Printf("-- After loop #[%d] - The Slice status is %v - Number of swaps done #[%d]\n", loopCount, aSliceToBeSorted, numberOfSwaps)

		if numberOfSwaps == 0 {
			isSorted = true
		}

		// infinite loop protection
		if loopCount > infiniteLoopProtection {
			fmt.Println("POTENTIAL INFINITE LOOP. The program will exit. ")
			os.Exit(1)
		}
	}
}

// Swap takes two arguments, a slice of integers and an index value i which indicates a position in the slice.
// it should swap the contents of the slice in position i with the contents in position i+1
func Swap(aSliceToBeSorted []int, position int) {
	aSliceToBeSorted[position], aSliceToBeSorted[position+1] = aSliceToBeSorted[position+1], aSliceToBeSorted[position]
}
