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
func main(){
nums:= []int{-1,0,3,5,9,12}
 target := 9
 search(nums,target)
}