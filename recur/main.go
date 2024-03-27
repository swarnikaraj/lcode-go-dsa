package main

import (
	"fmt"
)

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


// Recursive Programs to find Minimum and Maximum elements of array
// Given an array of integers arr, the task is to find the minimum and maximum element of that array using recursion.

func recurivemax(start int, arr []int, max int) int{
	if start==len(arr){
	return max
    }
  if arr[start]>max{
	max=arr[start]
  }
   return recurivemax(start+1,arr,max)
}
func findMax(arr []int) int{
 start:=0
 max:=arr[0]
sol:=recurivemax(start,arr,max)
return sol
}

// Program to check if an array is sorted or not (Iterative and Recursive)
// Given an array of size n, write a program to check if it is sorted in ascending order or not. Equal values are allowed in an array and two consecutive equal values are considered sorted.

func recursivecheck(start int, end int, arr []int) bool{
	if start==end{
		return true
	}

   if arr[start]<arr[start-1]{
	return false
   }
   return recursivecheck(start +1,end, arr)
}
func checkIfSortedArr(arr []int){
  start:=1
  end:=len(arr)
  sol:=recursivecheck(start, end, arr )
  fmt.Println(sol)
}


func  bubblesortIterative(arr []int){

 for i:=0;i<len(arr);i++{
    for j:=1;j<len(arr)-i;j++{
        if arr[j]<arr[j-1]{
            arr[j],arr[j-1]=arr[j-1],arr[j]
        }
    }
 }
 
  
}


func bubblesort(arr []int, i int, j int){
     if i==len(arr){
		return
	 }
    if j<len(arr)-i{

		if arr[j]<arr[j-1]{
            arr[j],arr[j-1]=arr[j-1],arr[j]
        }
      bubblesort(arr , i, j+1)
	}else{
     bubblesort(arr , i+1, 1)
	}   
}


func selectionsortIterative(arr []int){

	for i:=0;i<len(arr)-1;i++{
		min:=i
		for j:=i+1;j<len(arr);j++{
           if arr[j]<arr[min]{
			min=j
		   }
		}
		arr[i],arr[min]=arr[min],arr[i]
	}
}
func selectionsort(arr []int, i int , j int, max int){
   if i==0{
	return
   }
   if j<i{
	
	if arr[j]>arr[max]{
	 selectionsort(arr, i, j+1, j)
	}else{
	 selectionsort(arr, i, j+1, max)
	}
	 
   }else{
    arr[i-1],arr[max]=arr[max],arr[i-1]
	selectionsort(arr, i-1 , 0, 0)
   }
}

func mergeSort(){

}

func main() {
	// nums := []int{1, 4, 5, 3, 2}

	// selectionsort(nums,len(nums),0,0)
	// fmt.Println(nums)
// printTriangle(nums)
// fmt.Println(findMax(nums))
// checkIfSortedArr(nums)

}