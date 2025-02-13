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

	// Modifying an element ... time complexity : O(1), since if you know the index, then you can modify the element
	arr[2] = 6
	fmt.Printf("After modifying element at index 2, Array : %v\n", arr)

	// Insert an element at index 2 .. time complexity : O(n), because we need to shift all the remaining elements to the right of the position we intend to insert
	index := 2
	value := 7

	arr = append(arr[:index], append([]int{value}, arr[index:]...)...)
	fmt.Printf("After inserting 7 at index 2, Array : %v\n", arr)

	// Delete an element at index 2 .... time complexity : O(n), because it is a bit similar to actions performed during insertion
	index = 2
	arr = append(arr[:index], arr[index+1:]...)
	fmt.Printf("After deleting element at index 2, Array : %v\n", arr)

	fmt.Println(arr[len(arr)-1])

	// IMPORTANT:
	// create a slice of length 0 with the caapacity of arr
	numList := make([]int, 0, len(arr))
	//numList[0] = arr[0] // this is invalid
	numList = append(numList, arr[0]) // append is what can only be used to add elements to an array of length 0
	fmt.Println("The slice contains:", numList)

	// Note that to add or change the value of the first element in numList, the actions below is not valid
	// Lets say I want to change the elements of numList during a loop to the second element of arr
	for i := 0; i < len(arr); i++ {
		if i == 2 {
			fmt.Println("length of numList:", len(numList))
			fmt.Println("capacity of numList:", cap(numList))
			numList[i-1] = arr[i] // Invalid operation
		}
	}

}
