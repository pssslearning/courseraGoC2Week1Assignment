package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func swap(sli []int, pos int) {

	t := sli[pos]
	sli[pos] = sli[pos+1]
	sli[pos+1] = t

}

func BubbleSort(sli []int) {

	for i := 0; i < 10; i++ {
		for j := 0; j < 9; j++ {
			if sli[j] > sli[j+1] {
				swap(sli, j)
			}
		}
	}
}

func main() {

	a := make([]int, 10)
	fmt.Println("Enter 10 numbers to sort")

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 10; i++ {

		fmt.Println("Input number", i+1, " - ") // For formatting to break up each entry and responding print
		scanner.Scan()                          // Get the next value

		// Convert to an integer
		val, err := strconv.Atoi(scanner.Text())

		if nil == err {
			// Insert value into the array
			a[i] = val

		} else {
			// Print an error message
			fmt.Println("Value was not an integer")
			os.Exit(1)
		}

	}
	fmt.Println("UnSorted")
	fmt.Println(a)
	BubbleSort(a)
	fmt.Println("Sorted")
	fmt.Println(a)

}
