package main

import (
	"fmt"
	"sort"
)

func main() {

	array := []int{5, 4, 3, 2, 1}
	array = append(array, 7)

	x := 7

	fmt.Println(array)

	//Functions that search an element in an array
	sequentialSearch(x, array)
	binarySearch(x, array)
	//interpolationSearch(x, array)

}

//Sequential Search
func sequentialSearch(x int, arr []int) {

	i := 0
	for ; i < len(arr); i++ {
		if arr[i] == x {
			fmt.Printf("finded!, index: %d\n", i)
			break
		}
	}

	if i == len(arr) {
		fmt.Printf("Cannot find: %d\n", x)
	}

}

/*
Gets an array and split it through the process
*/
func binarySearch(x int, arr []int) int {

	//sort array
	sort.Ints(arr)

	indexMax := len(arr)

	fmt.Printf("Sorted: %v, arr lenght:%d\n", arr, len(arr))

	if len(arr) != 1 {

		middle := indexMax / 2
		fmt.Printf("middle:%d\n", middle)

		if arr[middle] == x {
			fmt.Printf("Finded!, value: %d\n", arr[middle])
			return arr[middle]
		}

		//if middle is greater than x
		if x < arr[middle] {
			return binarySearch(x, arr[0:middle])
		}

		//if middle is less than x
		if x > arr[middle] {
			return binarySearch(x, arr[middle:len(arr)])
		}

	}

	if arr[0] == x {
		fmt.Printf("Finded!, value: %d\n", arr[0])
		return arr[0]
	}

	fmt.Printf("Cannot find: %d\n", x)
	return -1

}

/*
Gets an array and split it through the process with intarpolation
*/
func interpolationSearch(x int, arr []int) int {

	//sort array
	sort.Ints(arr)

	fmt.Printf("Sorted: %v, arr lenght:%d\n", arr, len(arr))

	if len(arr) != 0 {

		middle :=
			((x - arr[0]) * (len(arr) - 0)) / (arr[len(arr)-1] -
				arr[0])

		fmt.Printf("middle:%d\n", middle)

		if arr[middle] == x {
			fmt.Printf("Finded!, value: %d\n", arr[middle])
			return arr[middle]
		}

		//if middle is greater than x
		if x < arr[middle] {
			return binarySearch(x, arr[0:middle])
		}

		//if middle is less than x
		if x > arr[middle] {
			return binarySearch(x, arr[middle:len(arr)])
		}

	}

	fmt.Printf("Cannot find: %d\n", x)
	return -1

}
