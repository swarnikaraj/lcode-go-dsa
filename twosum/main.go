package main

import "fmt"

// Question:Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]
// Example 3:
// Input: nums = [3,3], target = 6
// Output: [0,1]

var nums = []int{0,11,23,56,6}
var target = 6

// time complexity (o)n2
func twoSumN2(nums []int, target int) []int {
    var res=make([]int,2)

    for i:=0;i<len(nums);i++ {
        for j:=0;j<len(nums);j++ {
            if nums[i] +nums[j]==target && i!=j{
                res[0]=i;
                res[1]=j;
                return res
            }
        }
    }

    return res
}
// time complexity (o)n
func twoSumN1(nums []int, target int) []int {
    var mymap=make(map[int]int)
   for index,num :=range nums{
         diff:=target-num

        if j,found:=mymap[diff] ; found{
             return []int{index,j}
        }

       mymap[num]=index
            

   }

   
   
    return []int{-1,-1}
}
func main(){

	fmt.Print("Two Sum Problem")

	twoSumN2(nums, target)
    twoSumN1(nums,target)
	// fmt.Printf("your n2 solution %v", res)
}