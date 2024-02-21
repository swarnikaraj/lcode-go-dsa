package main

// Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

// You must write an algorithm with O(log n) runtime complexity.
func search(nums []int, target int) int {

    s:=0
    e:=len(nums)-1

    

    for s<=e{
        m:=s + (e-s)/2
        if nums[m]==target{
            return m
        }else if(nums[m]>target){
            e=m-1
        }else{
            s=m+1
        }


    }

    return -1
    
}
func main(){

}