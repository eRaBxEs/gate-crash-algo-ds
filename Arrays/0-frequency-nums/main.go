// You can edit this code!
// Click here and start typing.

/*
2. Return all the frequency of unique numbers from a given slice/array arr
from the smallest to the largest element of the given array/slice arr as an array/slice.
So if given an array/slice of [8, 3, 1, 2, 8, 3, 1, 1, 4, 8, 3]
The function should return [3, 1, 3, 1, 3]. Where in arr element 1 occurs 3 times,
2 occurs 1 time, 3 occurs 3 times, 4 occurs 1 time and 8 occurs 3 times.
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	//sliecInt := []int{4, 6, 7, 1, 9, 2, 5, 1, 6, 3, 4, 6, 7, 1, 2, 4, 3, 5, 6, 7, 7}
	sliceInt := []int{8, 3, 1, 2, 8, 3, 1, 1, 4, 8, 3}

	retV := uniqueCount(sliceInt)
	fmt.Println(retV)
}

func uniqueCount(numList []int) []int {
	sort.Ints(numList)
	uniqueNumCounts := make([]int, 0, len(numList)) // creating a slice of length 0 and a capacity of numList

	count := 1
	for i, _ := range numList {
		if i > 0 {
			if numList[i] == numList[i-1] {
				count++
				continue
			} else {

				uniqueNumCounts = append(uniqueNumCounts, count)
				count = 1

			}

		}
	}

	// append the last element's frequency once you are outside the loop
	uniqueNumCounts = append(uniqueNumCounts, count)

	return uniqueNumCounts

}
