package main

import "fmt"

// Sum triangle from array
// Given an array of integers, print a sum triangle from it such that the first level has all array elements. From then, at each level number of elements is one less than the previous level and elements at the level is be the Sum of consecutive two elements in the previous level.
// Example :

// Input : A = {1, 2, 3, 4, 5}
// Output : [48]
//          [20, 28]
//          [8, 12, 16]
//          [3, 5, 7, 9]
//          [1, 2, 3, 4, 5]

// Explanation :
// Here,   [48]
//         [20, 28] -->(20 + 28 = 48)
//         [8, 12, 16] -->(8 + 12 = 20, 12 + 16 = 28)
//         [3, 5, 7, 9] -->(3 + 5 = 8, 5 + 7 = 12, 7 + 9 = 16)
//         [1, 2, 3, 4, 5] -->(1 + 2 = 3, 2 + 3 = 5, 3 + 4 = 7, 4 + 5 = 9)

func printTriangle(arr []int) {
	if len(arr) <= 1 {
		fmt.Println(arr)
		return
	}
	output := []int{}

	for i := 1; i < len(arr); i++ {
		sum := arr[i] + arr[i-1]
		output = append(output, sum)
	}
	
	printTriangle(output)
	fmt.Println(arr)
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
printTriangle(nums)
}