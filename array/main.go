package main

import "fmt"

// Problem statement
// Given an array ‘arr’ of size ‘n’ find the largest element in the array.
// Example:

// Input: 'n' = 5, 'arr' = [1, 2, 3, 4, 5]

// Output: 5

// Explanation: From the array {1, 2, 3, 4, 5}, the largest element is 5.

func findlargestNumber(n int ,arr []int) int{
    largest:=arr[0]
	for _, num :=range arr{
		
		if(num>largest){
			largest=num
		}

	}
return largest
}

func main()  {
arr:=[]int{4 ,7, 8, 6 ,7 ,6 }
largestitem:=findlargestNumber(5,arr)
fmt.Print(largestitem," largestitem")

	
}

