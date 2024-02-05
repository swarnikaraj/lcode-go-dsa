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

func findSecLargestAndSecondSmallest(arr []int){

	for i:=0;i<2;i++{
		for j:=1;j<len(arr)-i;j++{
          if(arr[j]<arr[j-1]){
			arr[j], arr[j-1]=arr[j-1],arr[j]
		  }     
		}
	}
secondlargest:=arr[len(arr)-2]
	for i:=0;i<2;i++{
		for j:=1;j<len(arr)-i;j++{
          if(arr[j]>arr[j-1]){
			arr[j], arr[j-1]=arr[j-1],arr[j]
		  }     
		}
	}
secondsmallest:=arr[1]
	
fmt.Println(secondsmallest," secondsmallest sorted from bubble")
	fmt.Println(secondlargest," secondlargest sorted from bubble")
}

func main()  {
arr:=[]int{1 ,2, 8, 6 ,7 ,6 }
largestitem:=findlargestNumber(5,arr)
fmt.Println(largestitem," largestitem")
findSecLargestAndSecondSmallest(arr)
	
}

