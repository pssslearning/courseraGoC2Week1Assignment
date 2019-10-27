# courseraGoC2Week1Assignment

## Functions, Methods, and Interfaces in Go Week 1 Activity: Bubble Sort Program

Tareas calificadas por los compañeros: Module 1 Activity: Bubble Sort Program

- [Course: Functions, Methods, and Interfaces in Go](https://www.coursera.org/learn/golang-functions-methods)

- Course 2 of 3 from `Programming with Google Go Specialization`

## Instructions

```
The goal of this assignment is to create a routine that sorts a series of ten integers using the bubble sort algorithm. Click the "My Submission" tab to view specific instructions for this activity.

Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers. The program should print the integers out on one line, in sorted order, from least to greatest. Use your favorite search tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice. You should write a Swap() function which performs this operation. Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice. The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.

Bubble Sort is the simplest sorting algorithm that works by repeatedly swapping the adjacent elements if they are in wrong order.

```

## My assignment

Source code at `bubblesort/bubblesort.go`

## A sample test compilation and execution

```
... Test the program by running it twice and testing it with a different sequence of integers each time. The first test sequence of integers should be all positive numbers and the second test should have at least one negative number.
...
```

### Compiling
```sh
devel1@vbxdeb10mate:~$ cd $GOPATH/src/github.com/pssslearning/courseraGoC2Week1Assignment/bubblesort/
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC2Week1Assignment/bubblesort$ go build bubblesort.go 
```

### First test: All positive numbers
```sh
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC2Week1Assignment/bubblesort$ ./bubblesort 
========================================================
 You are now being asked to introduce up to 10 integers 
========================================================
Please enter integer #1 of 10 --> 234
Please enter integer #2 of 10 --> 1
Please enter integer #3 of 10 --> 345
Please enter integer #4 of 10 --> 123456
Please enter integer #5 of 10 --> 3
Please enter integer #6 of 10 --> 57
Please enter integer #7 of 10 --> 3
Please enter integer #8 of 10 --> 678
Please enter integer #9 of 10 --> 20
Please enter integer #10 of 10 --> 200
So far the initial slice built is [234 1 345 123456 3 57 3 678 20 200]:

================================
Entering BubbleSort Function
================================
So far the initial state of the slice received is [234 1 345 123456 3 57 3 678 20 200]:
-- After loop #[1] - The Slice status is [1 234 345 3 57 3 678 20 200 123456] - Number of swaps done #[7]
-- After loop #[2] - The Slice status is [1 234 3 57 3 345 20 200 678 123456] - Number of swaps done #[5]
-- After loop #[3] - The Slice status is [1 3 57 3 234 20 200 345 678 123456] - Number of swaps done #[5]
-- After loop #[4] - The Slice status is [1 3 3 57 20 200 234 345 678 123456] - Number of swaps done #[3]
-- After loop #[5] - The Slice status is [1 3 3 20 57 200 234 345 678 123456] - Number of swaps done #[1]
-- After loop #[6] - The Slice status is [1 3 3 20 57 200 234 345 678 123456] - Number of swaps done #[0]

=============================================================================
The resulting Slice status is [1 3 3 20 57 200 234 345 678 123456]
=============================================================================

```

### Second test: Should have at least one negative number
```sh
devel1@vbxdeb10mate:~/gowkspc/src/github.com/pssslearning/courseraGoC2Week1Assignment/bubblesort$ ./bubblesort 
========================================================
 You are now being asked to introduce up to 10 integers 
========================================================
Please enter integer #1 of 10 --> 345
Please enter integer #2 of 10 --> -200
Please enter integer #3 of 10 --> 34
Please enter integer #4 of 10 --> 0
Please enter integer #5 of 10 --> -345
Please enter integer #6 of 10 --> 123456
Please enter integer #7 of 10 --> 200
Please enter integer #8 of 10 --> 45
Please enter integer #9 of 10 --> -2
Please enter integer #10 of 10 --> 100
So far the initial slice built is [345 -200 34 0 -345 123456 200 45 -2 100]:

================================
Entering BubbleSort Function
================================
So far the initial state of the slice received is [345 -200 34 0 -345 123456 200 45 -2 100]:
-- After loop #[1] - The Slice status is [-200 34 0 -345 345 200 45 -2 100 123456] - Number of swaps done #[8]
-- After loop #[2] - The Slice status is [-200 0 -345 34 200 45 -2 100 345 123456] - Number of swaps done #[6]
-- After loop #[3] - The Slice status is [-200 -345 0 34 45 -2 100 200 345 123456] - Number of swaps done #[4]
-- After loop #[4] - The Slice status is [-345 -200 0 34 -2 45 100 200 345 123456] - Number of swaps done #[2]
-- After loop #[5] - The Slice status is [-345 -200 0 -2 34 45 100 200 345 123456] - Number of swaps done #[1]
-- After loop #[6] - The Slice status is [-345 -200 -2 0 34 45 100 200 345 123456] - Number of swaps done #[1]
-- After loop #[7] - The Slice status is [-345 -200 -2 0 34 45 100 200 345 123456] - Number of swaps done #[0]

=============================================================================
The resulting Slice status is [-345 -200 -2 0 34 45 100 200 345 123456]
=============================================================================

```
