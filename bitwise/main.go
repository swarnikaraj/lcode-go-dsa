package main

// 136. Single Number
// Solved
// Easy
// Topics
// Companies
// Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

// You must implement a solution with a linear runtime complexity and use only constant extra space.

func singleNumber(nums []int) int {
unique:=0
   
   for i:=0;i<len(nums);i++{
     unique ^=nums[i]
   }

   return unique
}


func main(){
	nums:=[]int{1,2,2,3,3,4,4,5,5}
singleNumber(nums)
}