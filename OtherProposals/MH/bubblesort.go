package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	for {
		list := make([]int, 0, 10)
		fmt.Printf("Enter up to 10 integers separated by spaces: ")

		reader := bufio.NewReader(os.Stdin)
		inputString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("%v", err)
			break
		}

		s := strings.Fields(inputString)
		if len(s) > 10 {
			fmt.Println("Too many arguments. Try again.")
			continue
		}

		for _, str := range s {
			number, err := strconv.Atoi(str)
			if err == nil {
				list = append(list, number)
			}
		}

		BubbleSort(list)

		fmt.Println("Sorted result: ", list)
	}
}

func BubbleSort(list []int) {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-1; j++ {
			Swap(list, j)
		}
	}
}

func Swap(list []int, index int) {
	i := index
	j := index + 1

	if list[i] > list[j] {
		temp := list[i]
		list[i] = list[j]
		list[j] = temp
	}
}
