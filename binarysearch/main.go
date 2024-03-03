package main

import (
	"fmt"
	"math"
)

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


func searchInInfiniteArray(nums []int,target int){
// find start and end and then binary search

 start:=0
 end:=1

 for target>nums[end]{
    newstart:=end+1
    end=end + (end-start+1)*2
    start=newstart

 }
 /// take start and end and do binary search
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


func peakIndexInMountainArray(arr []int) int {

    start:=0;
    end:=len(arr)-1
// try to find max
// every point of time. we r tryign to find max. and at last we will be left with only 1 item so that would be the best possible answer
    for start<end{
        mid:=start+ (end-start)/2

        if arr[mid]>arr[mid+1]{
            // means u r in decreasing part of mountain
            end=mid
        }else if arr[mid]<=arr[mid+1]{
            // means u r in increasing part of mountain
            start=mid+1
        }
    }
    // at lasst all start, mid and end will be at one point. and condition breaks so u can return start, end or mid
   return start 
}

func findInMountainArray(target int, mountainArr []int) int {

    // question asks for smallest index means the number should be the part of increasing portion
    // first find the peak
    // then make start=0 and end=peak
    /// and search there

    start:=0
    end:=len(mountainArr)-1
    found:=false
    for start<end{
        mid:=start + (end-start)/2
        if mountainArr[mid]>=mountainArr[mid+1]{
            end=mid
        }else {
            start=mid+1
        }
    }

   peak:=start
fmt.Println(mountainArr[peak], "found peak")
   newstart:=0
   newend:=peak 
fmt.Println(mountainArr[newstart], mountainArr[newend], found, "212")

   for newstart<=newend{
        mid:=newstart + (newend-newstart)/2

       if mountainArr[mid]==target{
           found=true
            return mid
        }else if mountainArr[mid]>target{
            newend=mid-1
        }else{
            newstart=mid+1
        }

   }
    
   
   fmt.Println(mountainArr[newstart], mountainArr[newend], found,"229")

   newstart=peak
   newend=len(mountainArr)-1

   fmt.Println(mountainArr[newstart], mountainArr[newend], found,"234")

if !found{
for newstart<=newend{
        mid:=newstart + (newend-newstart)/2

       if mountainArr[mid]==target{
            return mid
        }else if mountainArr[mid]>target{
            newstart=mid+1
            
        }else{
            newend=mid-1
        }

   }
}
   

   return -1
}

func rotatedsearch(nums []int, target int) int {
  start:=0
end:=len(nums)-1

  bs:=func (nums []int, target int, start int, end int) int{

        for start<=end{
            mid:=start + (end-start)/2

            if nums[mid]==target{
                return mid
            }
           if nums[mid]>target{
          end=mid-1
            }else{
                start=mid+1
            }
        }
return -1
    }
    
  findpivot:= func(nums []int) int {
      start:=0
      end:=len(nums)-1

      for start<=end{
           mid:=start + (end-start)/2
          if mid<end && nums[mid]>nums[mid+1]{
              return mid
          }

          if mid>start && nums[mid]<nums[mid-1]{
              return mid-1
          }

          if nums[start]>nums[mid]{
              end=mid-1
          }else{
              start=mid+1
          }

      }
  return -1
  }
 
 pivot:=findpivot(nums)
 fmt.Println(pivot, "pivot hu")
if pivot !=-1{
   if nums[pivot]==target{
     return pivot
 }
 if nums[start]==target{
     return start
 }else if nums[start]>target{
     ans:=bs(nums,target,pivot+1,end)
     return ans
 }else{
     ans:=bs(nums,target,0, pivot)
     return ans
 }
 }
 fmt.Println("Pivot was", pivot)
  ans:=bs(nums,target,0,end)

return ans
}

func findPivot(nums []int)int{
    start:=0
    end:=len(nums)-1

    for start<=end{
        mid:=start + (end-start)/2

        if nums[mid]>nums[mid+1] && mid<end{
            return mid
        }
        if nums[mid]<nums[mid-1] && mid>start{
            return mid-1
        }
        if start>mid{
            end=mid-1
        }else{
            start=mid+1
        }
    }

    return -1
}

func searchMatrix(matrix [][]int, target int) bool {
   row:=0
   col:=len(matrix[0])-1

   for row<len(matrix) && col>=0{
       elm:=matrix[row][col]
       if target<elm{
           col--
       }else if target>elm{
           row++
       }else{
           return true
       }
   } 

   return false
}
func findMin(nums []int) int {
  start:=0
    end:=len(nums)-1

    for start<=end{
        mid:=start + (end-start)/2

        if mid<end && nums[mid] > nums[mid+1] {
            return nums[mid+1]
        }
        if mid>start && nums[mid]<nums[mid-1] {
            return nums[mid]
        }
        if nums[start]>nums[mid]{
            end=mid-1
        }else{
            start=mid+1
        }
    }

    return nums[0]
}

func singleNonDuplicate(nums []int) int {
   start := 0
    end := len(nums) - 1

    for start < end {
        mid := start + (end-start)/2

        // Ensure mid is even to maintain pairs
        if mid%2 == 1 {
            mid--
        }

        // Check if the single element is on the right or left side of mid
        if nums[mid] != nums[mid+1] {
            end = mid
        } else {
            start = mid + 2
        }
    }

    // At this point, left and right pointers converge to the single element
    return nums[start]
}

func searchInsoretedRepeated(nums []int, target int) int {
  start:=0
  end:=len(nums)-1

  bs:=func (nums []int, target int, start int, end int) int{

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
return -1
    }
    
  findpivot:= func(nums []int) int {
      start:=0
      end:=len(nums)-1

      for start<=end{
           mid:=start + (end-start)/2
           
          if mid<end && nums[mid]>nums[mid+1]{
              return mid
          }

          if mid>start && nums[mid]<nums[mid-1]{
              return mid-1
          }

         if nums[start]==nums[mid] && nums[mid]==nums[end]{


             if nums[start]>nums[start+1]{
                 return start
             }
             start++
             if nums[end]<nums[end-1]{
                 return end-1
             }
             end--
         }else if nums[start]>=nums[mid] && nums[start]>nums[end]{
             start=mid+1
         }else{
             end=end-1
         }

      }
  return -1
  }
 
 pivot:=findpivot(nums)
 fmt.Println(pivot, "pivot hu")
if pivot !=-1{
   if nums[pivot]==target{
     return pivot
 }
 if nums[start]==target{
     return start
 }else if nums[start]>target{
     ans:=bs(nums,target,pivot+1,end)
     return ans
 }else{
     ans:=bs(nums,target,0, pivot)
     return ans
 }
 }
 fmt.Println("Pivot was", pivot)
  ans:=bs(nums,target,0,end)

return ans
}


func squareroot( n int) int {
   start:=0
   end:=n

   for start<=end{

    mid:=start + (end-start)/2

    if mid*mid ==n{
        return mid
    }
    if mid*mid<n{
        start=mid+1
    }else{
        end=mid-1
    }
   }

   return -1
}

func expo(base int, power int){
    ans:=1

    for power>0{
        ans=ans*base
        power--
    }
}
func totalhours(hr int, piles []int) int {
    totalhour := 0
    for i := 0; i < len(piles); i++ {
        totalhour += int(math.Ceil(float64(piles[i]) / float64(hr)))
    }
    return totalhour
}

func minEatingSpeed(piles []int, h int) int {
    max := piles[0]
    for i := 1; i < len(piles); i++ {
        if piles[i] > max {
            max = piles[i]
        }
    }

    start := 1
    end := max

    for start < end {
        mid := start + (end - start) / 2
        totalhr := totalhours(mid, piles)
        if totalhr <= h {
            end = mid
        } else {
            start = mid + 1
        }
    }
    return start
}

func nthroot( n int, m int) int {
   start:=1
   end:=m

   for start<=end{

    mid:=start + (end-start)/2

    if m==int(math.Pow(float64(mid), float64(n))){
        return mid
    }
    if m<int(math.Pow(float64(mid), float64(n))){
        end=mid-1
       
    }else{
         start=mid+1
    }

   }

   return -1
}

func main(){
// nums:= []int{-1,0,3,5,9,12}
//  target := 9
//  search(nums,target)

// nums:=[]int{1,3}
// target:=2
// sol:=searchInsert(nums,target)
// fmt.Println(sol)
// target:=2
// nums:=[]int{1,5,2}
// sol:=findInMountainArray(target,nums)
// fmt.Println(sol)

// nums:=[]int{1,2,3,4,5,6,7,8}
// sol:=findPivot(nums)
// fmt.Println(sol)
// sol:=rotatedsearch(nums, 4)

// sol:=squareroot(25)

sol:=nthroot(3,64)
fmt.Println(sol)
}
