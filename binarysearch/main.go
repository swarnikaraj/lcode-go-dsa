package main

import "fmt"

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

// find first and last occurence
func searchRange(nums []int, target int) []int {

    res:=[]int{-1,-1}

    start:=0;
    end:=len(nums)-1
// firstoccurence
    for start<=end{
      mid:=start + (end-start)/2

      if nums[mid]==target{
          res[0]=mid
          end=mid-1
      }else if nums[mid]<target{
          start=mid+1
      }else{
          end=mid-1
      }

    }



    start=0
    end=len(nums)-1
    
     for start<=end{
      mid:=start + (end-start)/2

      if nums[mid]==target{
          res[1]=mid
          start=mid+1
      }else if nums[mid]<target{
          start=mid+1
      }else{
          end=mid-1
      }

    }

   

    return res
}

func searchInsert(nums []int, target int) int {

    start:=0;
    end:=len(nums)-1
        if target<=nums[0]{
            return 0
        }
        if target>nums[len(nums)-1]{
            return len(nums)
        }
    for start<=end{
        
        mid:=start + (end-start)/2
   
        if nums[mid]==target{
            return mid
        }else if nums[mid]>target{
            end=mid-1
        }else{
            start=mid+1
        }
    }
    return start
}

// greatest number smaller than target
func floorthecel(nums []int, target int) int {

    start:=0;
    end:=len(nums)-1
        if target<=nums[0]{
            return 0
        }
        if target>nums[len(nums)-1]{
            return len(nums)-1
        }
    for start<=end{
        
        mid:=start + (end-start)/2
  
        if nums[mid]==target{
            return mid
        }else if nums[mid]>target{
            end=mid-1
        }else{
            start=mid+1
        }
    }
    return end
}
// smallest number greater than target
func ceilof(nums []int, target int) int {
    start := 0
    end := len(nums) - 1

    if target < nums[0] {
        return 0 // The first element is the smallest item greater than the target
    }
    if target >= nums[len(nums)-1] {
        return -1 // There's no element greater than the target, return -1 or handle it as appropriate in your context
    }

    for start <= end {
        mid := start + (end-start)/2

        if nums[mid] == target {
            return mid + 1 // The next element is the smallest item greater than the target
        } else if nums[mid] < target {
            start = mid + 1 // Move to the right half
        } else {
            end = mid - 1 // Move to the left half
        }
    }
    return start // The element at 'start' will be the smallest item greater than the target
}
func main(){
// nums:= []int{-1,0,3,5,9,12}
//  target := 9
//  search(nums,target)

nums:=[]int{1,3}
target:=2
sol:=searchInsert(nums,target)
fmt.Println(sol)
}