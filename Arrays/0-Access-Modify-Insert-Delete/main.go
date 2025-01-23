package main

import (
	"fmt"
)

func main() {
	// Declaration and initialization of slice:array
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("Printing the array: %v\n", arr)
	fmt.Printf("Printing more details of the given array: %#v\n", arr)

	// Accessing an element ... Time complexity : O(1), because if you know the index you will get the element
	fmt.Printf("Element at index 2 is %v\n", arr[2])

	// Accessing the elements of an array using for loop ... Time complexity : O(n) because if you have n elements the loop will run n times to access each element
	fmt.Printf("Accessing all elements in %v:\n", arr)
	for x, y := range arr {
		fmt.Printf("Element at index %d is %v\n", x, y)
	}

}
