package main

/*

1. Print/return all the unique numbers from the smallest to the largest, given a function below, where numList is a list of numbers [2, 1, 9, 2, 4, 8, 1, 9, 4, 5, 2, 6, 7, 10, 8, 3, 4, 11, 15, 19] and each element can occur as many times as possible
func unique(numList []int) []int {
	// implement here
}
I have implemented them in Go below:
https://go.dev/play/p/Z45GR1l4_r-
https://go.dev/play/p/InO-ozc6F9G
https://go.dev/play/p/n1_uK51igPZ

*/
import (
	"fmt"
	"sort"
)

func unique(numList []int) []int {
	sort.Ints(numList)
	uniqueNums := make([]int, 0, len(numList)) // creating a slice of length 0 and a capacity of numList

	for i, v := range numList {
		if i == 0 || numList[i] != numList[i-1] {
			uniqueNums = append(uniqueNums, v)
		}
	}
	return uniqueNums
}

func main() {
	sliecInt := []int{2, 1, 9, 2, 4, 8, 1, 9, 4, 5, 2, 6, 7, 10, 8, 3, 4, 11, 15, 19}

	// Tasks highlighted below
	//1. Print out all the unique numbers from the smallest to the largest

	retV := unique(sliecInt)
	fmt.Println(retV)
}
