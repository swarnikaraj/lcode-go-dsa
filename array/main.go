package main

import (
	"fmt"
	"math"
)

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

func findSecLargestAndSecondSmallest(n int,arr []int) (int, int){
 max:=arr[0]
 min:=arr[0]

  for i:=0;i<len(arr);i++{
     if arr[i]>max{max=arr[i]}
	 if arr[i] <min {min=arr[i]}
  }
  fmt.Print(max,"Max")
  fmt.Print(min,"min")
    secondLargest:=math.MinInt64
	secondSmallest:=math.MaxInt64
  for j:=0;j<len(arr);j++{
     if arr[j]!=max && arr[j]>secondLargest {
		secondLargest=arr[j]
	 }
	  if arr[j]<secondSmallest && arr[j]!=min{
		secondSmallest=arr[j]
	 }
  }

return secondLargest, secondSmallest
}

func checkifArrayIssorted(n int, arr []int) bool{

	res:=true

	for i:=1;i<len(arr);i++{
		if arr[i-1]>arr[i]{
			res=false

		}
	}
 return res
}

func isArrySortedAndRotated(n int, nums []int)bool{
	rotpoint:=0; res:=true
	if(nums[0]>=nums[len(nums)-1]){
		
     for i:=1;i<len(nums);i++{
    if nums[i]<nums[i-1]{
    rotpoint=i
	break
	}
      }
	}


      for j:=rotpoint+1;j<len(nums);j++{
          if nums[j]<nums[j-1]  {
			res=false
		  }
	  }
	

   return  res
}

func removeDuplicates(nums []int) int {

    i:=0
    j:=1

    for  j<len(nums) && i<len(nums)-1{
        
      if nums[j]!=nums[i]{
          nums[i+1]=nums[j]
          i++
      }
      j++
    }
   return i+1
}
func main()  {
arr:=[]int{1 ,2, 8, 6 ,7 ,6 }
largestitem:=findlargestNumber(5,arr)
fmt.Println(largestitem," largestitem")
sL, sm:=findSecLargestAndSecondSmallest(5,arr)
fmt.Printf("second largest %v and second smallest %v",sL, sm)
	issorted:=checkifArrayIssorted(5,arr)
	fmt.Print(issorted," sorted result")
	nums := []int{1, 2, 2, 1}
	issortandrot:=isArrySortedAndRotated(4,nums)
	fmt.Println(issortandrot,"Is sorted and rotated")
	dpres:=removeDuplicates(nums)
	fmt.Println(dpres,"removed duplicate")
}

