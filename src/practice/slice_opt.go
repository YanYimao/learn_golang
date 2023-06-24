package main

import "fmt"

type MySlice []int

func (slice *MySlice) add(index int, num int) {
	*slice = append((*slice)[:index], append([]int{num}, (*slice)[index:]...)...)
}

func (slice *MySlice) remove(index int) {
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
}

func main() {
	slice := MySlice{1, 2, 3, 4}
	fmt.Printf("slice: %v\n", slice)

	slice.add(0, 0)
	fmt.Printf("slice: %v\n", slice)

	slice.remove(4)
	fmt.Printf("slice: %v\n", slice)
}
