package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Swap i and i+1 values in vv
func Swap(vv []int, i int) {
	temp := vv[i]
	vv[i] = vv[i+1]
	vv[i+1] = temp
}

//BubbleSort (destructive)
func BubbleSort(vv []int) {
	for i := range vv {
		for j := range vv[:len(vv)-i-1] {
			if vv[j] > vv[j+1] {
				Swap(vv, j)
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter up to 10 integers separated by spaces:")
	ui, _ := reader.ReadString('\n')
	ui = strings.TrimSuffix(ui, "\n")
	ui = strings.TrimSuffix(ui, "\r")

	ss := strings.Split(ui, " ")
	if len(ss) > 10 {
		ss = ss[:10]
	}
	vv := make([]int, 0, len(ss))
	for _, s := range ss {
		v, err := strconv.Atoi(s)
		if err == nil {
			vv = append(vv, v)
		}
	}

	BubbleSort(vv)

	fmt.Println(vv)
}
