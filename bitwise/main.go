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
// 338. Counting Bits
// Solved
// Easy
// Topics
// Companies
// Hint
// Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.


func count1(n int) int{
    sum:=0
    for n>0{
        last:=n & 1
        sum=sum+last
        n=n>>1
    }
    return sum
}
func countBits(n int) []int {
    ans:=make([]int,n+1)

    i:=1

    for i<n+1{
       
       total1:=count1(i)
       ans[i]=total1
    

        i++
    }
   return ans 
}
func main(){
	nums:=[]int{1,2,2,3,3,4,4,5,5}
singleNumber(nums)
countBits(9)
}