package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strings"
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

func removeDuplicates2(nums []int) int {

   obj:=make(map[int]int)

   for i:=0;i<len(nums);i++{
       obj[nums[i]]=nums[i]
   }

   arr:=[]int{}
   for _, num:=range obj{
       arr=append(arr,num)
   }
 sort.Ints(arr)
  for index,ar:=range arr{
   nums[index]=ar
  }

  return len(arr)
}

func singleNumber(nums []int) int {

    obj:=make(map[int]int)

    for _, num :=range nums{
		 value, ok := obj[num]
        if ok {
             obj[num]= value+1
        }else{
             obj[num]=1
        }
    }

    for key,value :=range obj{
       if value==1{
		return key
	   }
        
    }
    return -1
}

func moveZeroes(nums []int)  {
   nonzeros := make([]int, len(nums))
    index:=0
    zeros:=0
    for i:=0;i<len(nums);i++{
      if nums[i]!=0{
          nonzeros[index]=nums[i]
          index++
      }else{
          zeros++
      }
    }

    for i:=0;i<index;i++{
       nums[i] =nonzeros[i]
    }

    for i:=index;i<len(nums);i++{
        nums[i]=0
    }
}

// Problem statement
// Given two sorted arrays, ‘a’ and ‘b’, of size ‘n’ and ‘m’, respectively, return the union of the arrays.
// The union of two sorted arrays can be defined as an array consisting of the common and the distinct elements of the two arrays. The final array should be sorted in ascending order.
// Note: 'a' and 'b' may contain duplicate elements, but the union array must contain unique elements.
// Example:
// Input: ‘n’ = 5 ‘m’ = 3


func uniosOfsorted(a []int, b []int) []int{
  n:=len(a)
  i:=0
  m:=len(b)
  j:=0
  newarray:=make([]int,0)
  k:=0
 
   for i<n && j<m {
	
	if a[i]<b[j]{
		newarray = append(newarray, a[i])
     
	  i++
	  
	}else if a[i]==b[j]{
		newarray = append(newarray, a[i])
		 i++
		 j++
	}else{
		newarray = append(newarray, b[j])
	  j++
	}

	k++
  }
   
  for i<len(a){
	newarray = append(newarray, a[i])
	i++
  }

   for j<len(b){
	newarray = append(newarray, b[j])
	j++
   }

  return newarray
}

// Problem statement
// You are given an array 'a' of size 'n' and an integer 'k'.
// Find the length of the longest subarray of 'a' whose sum is equal to 'k'.
// Example :
// Input: ‘n’ = 7 ‘k’ = 3
// ‘a’ = [1, 2, 3, 1, 1, 1, 1]

// Output: 3

// Explanation: Subarrays whose sum = ‘3’ are:
// [1, 2], [3], [1, 1, 1] and [1, 1, 1]
// Here, the length of the longest subarray is 3, which is our final answer.

func printsubarray(arr []int){
for i:=0;i<len(arr);i++{

	for j:=i+1;j<len(arr)+1;j++{
		fmt.Println(arr[i:j])
	}
}
}

func longestSubarrayWithSumK(arr []int, K int) (int){
	longest:=0
    
	
   for i:=0;i<len(arr);i++{
	 for j:=i+1;j<len(arr);j++{
		
       subarray:=arr[i:j]
	   sum:=0
       for k:=0;k<len(subarray);k++{
         sum=sum+subarray[k]
	   }
	   if sum==K && longest<len(subarray){
		longest=len(subarray)
		
	   }
	}
   }

   return  longest
}

func rotateRight(nums []int, k int)  {
     rot:=k% int(len(nums))
	 
    for i:=1;i<=rot;i++{
		temp:=nums[len(nums)-1]
        for j:=len(nums)-2;j>=0;j--{
            nums[j+1]=nums[j]
        }
        nums[0]=temp
	}
	fmt.Print(nums)
}
func rotateLeft(nums []int, k int)  {
    rot:=k% int(len(nums))
	 
     for i:=1;i<=rot;i++{
        temp:=nums[0]
		for j:=1;j<len(nums);j++{
           nums[j-1]=nums[j]
		}
         nums[len(nums)-1]=temp
	 }
  fmt.Print(nums)
	 
}
// [1,2,3,4,5,6,7]
func OptimizerotateRight(nums []int, k int)  {
     k=k% int(len(nums))
	 temp := []int{}
     for i:=k;i>=1;i--{
          temp=append(temp, nums[len(nums)-i])
	 }

	 for i:=0;i<len(nums)-k;i++{
		temp=append(temp, nums[i])
	 }
   for i:=0;i<len(nums);i++{
	nums[i]=temp[i]
     }
	fmt.Print(nums)
}
// Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.
// We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
// You must solve this problem without using the library's sort function.
// Example 1:

// Input: nums = [2,0,2,1,1,0]
// Output: [0,0,1,1,2,2]

func sortColors(nums []int)  {
    twos:=0
    ones:=0
    zeros:=0;
fmt.Println(nums)
    for _, num :=range nums{
        if num==0{
            zeros++
        }else if num==1{
           ones++
        }else{
            twos++
        }
    }

	index:=0


	for(index<zeros){
nums[index]=0
index++
	}

	for index1:=0;index1<ones; index1++{
nums[index]=1
index++

	}
    
	for index2:=0;index2<twos; index2++{
nums[index]=2
index++

	}
	
   
}


// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.


// Example 1:

// Input: nums = [3,2,3]
// Output: 3


func majorityElement(nums []int) int {
    obj:=make(map[int]int)


    for _, num :=range nums{
        if _, exist:=obj[num] ; exist{
			
             obj[num]=obj[num]+1
        }else{
            obj[num]=1
        }
    }

	
  for key , value :=range obj{
	if value>len(nums)/2{
      return key
	}
  }

  return 0
}

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

func twoSum(nums []int, target int) []int {
  obj:=make(map[int]int)

  for index, num:=range nums{
         diff:=target-num
      if j , exist:= obj[diff] ; exist{
			return []int{index , j}
      }

      obj[num]=index
  }

   
   
    return []int{-1,-1}
}


// 53. Maximum Subarray
// Medium
// Topics
// Companies
// Given an integer array nums, find the 
// subarray
//  with the largest sum, and return its sum.
// Example 1:
// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: The subarray [4,-1,2,1] has the largest sum 6.
func arrSum(subarr []int) int{
	sum:=0
 for k:=0;k<len(subarr);k++{
            sum=sum+subarr[k]
		  }
		  return sum
}

func maxSubArray(nums []int)  int{
	max:=math.MinInt
	
    for start:=0;start<len(nums);start++{
		 
             
		for end:=start;end<len(nums);end++{
             sum:=0
			 for i:=start;i<=end;i++{
               sum=sum+nums[i]
			 }
		  if(max<sum){
			fmt.Println("Yes I am smaller")
		     max=sum
		  }

			
		 }
	}
   fmt.Println(max," mai zero kaise")
	return max
}

func isPalindrome(s string) bool {

   reg := regexp.MustCompile("[^a-zA-Z0-9]+")


    st:=reg.ReplaceAllString(s,"")
    str:=strings.ToLower(st)

    fmt.Print(str)

    i:=0
    j:=len(str)-1

    for i<j{
        if str[i]!=str[j]{
            return false
        }
        i++
        j--
    }

    return true
}
func isSubsequence(s string, t string) bool {
    i:=0
    j:=0

    for i<len(s) && j<len(t){
        if s[i]==t[j]{
            i++
        }

        j++
    }

    return i==len(s)
}
func findDuplicate(nums []int) int {
  res:=0
    numcount := make(map[int]int)
    for _, nm := range nums {
        numcount[nm]++
    }
   for key, value:=range numcount{
      if value!=1{
          res=key
          break
      }
   }
    fmt.Println(numcount)

   
    return res
}
func kadanesAlgo(nums []int) int{
	  maxEndingHere := nums[0]
    maxSoFar := nums[0]

    for i := 1; i < len(nums); i++ {
        maxEndingHere = int( math.Max( float64(nums[i]), float64(maxEndingHere+nums[i]) )  )
        maxSoFar = int(math.Max(float64(maxSoFar), float64(maxEndingHere)))
    }

    return maxSoFar

		
}

// 121. Best Time to Buy and Sell Stock
// Solved
// Easy
// Topics
// Companies
// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
func maxProfit(prices []int) int {
    maxprofit:=0
    min:=prices[0]
    for i:=1 ;i<len(prices);i++{
        min=int(math.Min(float64(prices[i-1]),float64(min)))
        maxprofit=int(math.Max(float64(prices[i]-min),float64(maxprofit)))
    }
   return maxprofit 
}

// 2149. Rearrange Array Elements by Sign
// Medium
// Topics
// Companies
// Hint
// You are given a 0-indexed integer array nums of even length consisting of an equal number of positive and negative integers.

// You should rearrange the elements of nums such that the modified array follows the given conditions:

// Every consecutive pair of integers have opposite signs.
// For all integers with the same sign, the order in which they were present in nums is preserved.
// The rearranged array begins with a positive integer.
// Return the modified array after rearranging the elements to satisfy the aforementioned conditions.

func rearrangeArray(nums []int) []int {
    pos:=[]int{}
	
	neg:=[]int{}

	for i:=0;i<len(nums);i++{
		if nums[i]>=0{
			pos=append(pos, nums[i])
		}else{
			neg=append(neg,nums[i] )
		}
	}

	
    i:=0
    j:=0
	for k:=0;k<len(nums) ;k++{
		if k%2==0 || k==0{
			nums[k]=pos[i]
            i++
		}else{
			nums[k]=neg[j]
            j++
		}
	}

 return nums
}
func rearrangeArray2(nums []int) []int {
    pos:=[]int{}
	
	neg:=[]int{}

	for i:=0;i<len(nums);i++{
		if nums[i]>=0{
			pos=append(pos, nums[i])
		}else{
			neg=append(neg,nums[i] )
		}
	}

	
    
	for k:=0;k<len(nums)/2 ;k++{
		
			nums[2*k]=pos[k]
           
	
			nums[2*k+1]=neg[k]
           
		
	}

 return nums
}
func rearrangeArray3(nums []int) []int {
  
	ans:=make([]int,len(nums))
	pos:=0; neg:=1

	for i:=0;i<len(nums);i++{
		if nums[i]>=0{
			ans[pos]=nums[i]
            pos=pos+2
		}else{
			ans[neg]=nums[i]
            neg=neg+2
		}
	}

	
    
	

 return ans
}

// 54. Spiral Matrix
// Medium
// Topics
// Companies
// Hint
// Given an m x n matrix, return all elements of the matrix in spiral order.

func spiralOrder(matrix [][]int) []int {
    ans:=[]int{}
	m:=len(matrix)
	n:=len(matrix[0])
   left:=0;top:=0;bot:=m-1; right:=n-1
   count:=0
   for count<m*n{

	for i:=left;i<=right && count<m*n;i++{
		ans=append(ans,matrix[top][i])
		count++
       
	}
	top++

	for j:=top;j<=bot  && count<m*n;j++{
		ans = append(ans, matrix[j][right])
		count++
	}
	right--

	for k:=right;k>=left  && count<m*n;k--{
		ans = append(ans, matrix[bot][k])
		count++
	}
	bot--

	for l:=bot;l>=top  && count<m*n;l--{
		ans = append(ans, matrix[l][left])
		count++
	}
	left++

   }

   
  fmt.Print(ans)
	return ans
}


// 48.Rotate Image
// Medium
// Topics
// Companies
// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

func rotate90(matrix [][]int)  {
     
    for i:=0;i<len(matrix)-1;i++{
      for j:=i+1;j<len(matrix[0]);j++{
        matrix[i][j],matrix[j][i]=matrix[j][i],matrix[i][j]
      }
    }
  
    for i:=0;i<len(matrix);i++{
      row:=matrix[i]
       index:=0; k:=len(row)-1
     

      for index<=k{
        row[index],row[k]=row[k],row[index]
        index++
        k--
      }

    }
  fmt.Print(matrix)
}

// 560. Subarray Sum Equals K
// Medium
// Topics
// Companies
// Hint
// Given an array of integers nums and an integer k, return the total number of subarrays whose sum equals to k.

// A subarray is a contiguous non-empty sequence of elements within an array.

func subarraySum(nums []int, k int) int {
    sum:=0; count:=0
    for i:=0;i<len(nums);i++{
      sum=0;
      for j:=i;j<len(nums);j++{
     
       sum=sum+nums[j]
       fmt.Println(sum, " ", nums[i])
       if sum==k{
        fmt.Println(sum, "ho gya sum")
           count++
           sum=0
       }
      }
      
    }
    return count
}
 
//prefixsum
func subarraySum2(nums []int, k int) int {
    prefixsum:=0; count:=0
    myobj:=make(map[int]int)
    myobj[0]=1

    for i:=0;i<len(nums);i++{
     prefixsum=prefixsum+nums[i]
    
         toexclude:=prefixsum-k
     
           count=count+myobj[toexclude]
   
         myobj[prefixsum]=myobj[prefixsum]+1
     
    
    }
    return count
}
 



// 73. Set Matrix Zeroes
// Medium
// Topics
// Companies
// Hint
// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
// You must do it in place.

func setZeroes(matrix [][]int)  {
 
     for i:=0;i<len(matrix);i++{
      for j:=0; j<len(matrix[0]);j++{

        if matrix[i][j]==0{
           r:=i
           c:=j
         
             for i:=0;i<len(matrix);i++{
              if matrix[i][c] !=0{
                matrix[i][c]=math.MaxInt
              }
             }

             for i:=0;i<len(matrix[0]);i++{
              if matrix[r][i] !=0{
                matrix[r][i]=math.MaxInt
              }
             }
          
        }
      }
     }

   
       for i:=0;i<len(matrix);i++{
       for j:=0; j<len(matrix[0]);j++{
        if matrix[i][j]==math.MaxInt{
         matrix[i][j]=0
        }
      }
     }
     
}

func setZeroes2(matrix [][]int)  {
   rows:=make([]int, len(matrix))
   cols:=make([]int, len(matrix[0]))
     for i:=0;i<len(matrix);i++{
      for j:=0; j<len(matrix[0]);j++{

        if matrix[i][j]==0{
          rows[i]=1
          cols[j]=1
        }
      }
     }

   
       for i:=0;i<len(matrix);i++{
       for j:=0; j<len(matrix[0]);j++{
        if rows[i]==1 || cols[j]==1{
         matrix[i][j]=0
        }
      }
     }

     fmt.Println(matrix)
     
}

// Problem statement
// There is an integer array ‘a’ of size ‘n’.
// An element is called a Superior Element if it is greater than all the elements present to its right.
// You must return an array all Superior Elements in the array ‘a’.
// Note:
// The last element of the array is always a Superior Element. 
// Example
// Input: a = [1, 2, 3, 2], n = 4
// Output: 2 3
// Explanation: 
// a[ 2 ] = 3 is greater than a[ 3 ]. Hence it is a Superior Element. 
// a[ 3 ] = 2 is the last element. Hence it is a Superior Element.
// The final answer is in sorted order.

func leaderItems(arr []int){
  
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
   res:=[]int{}    
   i:=0 
    j:= 0
   for i<len(nums1) && j<len(nums2){
       if nums1[i]<=nums2[j]{
           res=append(res,nums1[i])
           i++
       }else{
           res=append(res,nums2[j])
           j++ 
       }
   }
  
  for i<len(nums1){
      res=append(res,nums1[i])
      i++
  }

   for j<len(nums2){
      res=append(res,nums2[j])
     j++
  }

 n:=len(res)
 ans:=0.0000
 if n%2==1 {
    ans= float64(res[n/2])
 }else{
    ans= (float64(res[n/2])+float64(res[(n/2)-1]))/2.0000
 }
  return ans  
}

func firstMissingPositive(nums []int) int {
  
    largest:=0
       myobj:=make(map[int]int)
    for i:=0;i<len(nums);i++{
        if nums[i]>largest{
            largest=nums[i]
        }
        if nums[i]>0{
           myobj[nums[i]]=nums[i]
        }
       
    }

    fmt.Println(myobj, " object")

    fmt.Println(largest, "  largest")

    for j:=1;j<=largest;j++{
     
      ans, exist:=myobj[j]
      fmt.Println(ans,exist)
      if(!exist){
        return j
      }
     
    }

return largest+1
}

func digitcal(digit int) int{
       c:=0
       for digit>0{
           digit=digit/10
           c++       }
fmt.Print(c,"couted digi")
           return c
   }
func findNumbers(nums []int) int {
    count:=0

   

    for i:=0;i<len(nums);i++{
       dc:=digitcal(nums[i])

       if dc%2==0{
           count++
       }
    }
    return count
}

// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

// Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:

// Change the array nums such that the first k elements of nums contain the elements which are not equal to val. The remaining elements of nums are not important as well as the size of nums.
// Return k.
// Custom Judge:

// The judge will test your solution with the following code:

// int[] nums = [...]; // Input array
// int val = ...; // Value to remove
// int[] expectedNums = [...]; // The expected answer with correct length.
//                             // It is sorted with no values equaling val.

// int k = removeElement(nums, val); // Calls your implementation

// assert k == expectedNums.length;
// sort(nums, 0, k); // Sort the first k elements of nums
// for (int i = 0; i < actualLength; i++) {
//     assert nums[i] == expectedNums[i];
// }
// If all assertions pass, then your solution will be accepted.

 

// Example 1:

// Input: nums = [3,2,2,3], val = 3
// Output: 2, nums = [2,2,_,_]
// Explanation: Your function should return k = 2, with the first two elements of nums being 2.
// It does not matter what you leave beyond the returned k (hence they are underscores).
// Example 2:

// Input: nums = [0,1,2,2,3,0,4,2], val = 2
// Output: 5, nums = [0,1,4,0,3,_,_,_]
// Explanation: Your function should return k = 5, with the first five elements of nums containing 0, 0, 1, 3, and 4.
// Note that the five elements can be returned in any order.
// It does not matter what you leave beyond the returned k (hence they are underscores).

func removeElement(nums []int, val int) int {

i:=0
    for i<len(nums){

        if nums[i]==val{
            j:=i+1
            for j<len(nums){
                if nums[j]!=val{
                nums[i], nums[j]=nums[j], nums[i]
                break
                }
                j++
            }
           if j==len(nums){
           return i
          
           }

        }
          i++
        
    }


   
    return len(nums)
}

func containsNearbyDuplicate(nums []int, k int) bool {
    obj:=make(map[int]int)

    for i:=0;i<len(nums);i++{
        if mapp, ok:=obj[nums[i]];ok{
           if int(math.Abs(float64(i)- float64(mapp)))>=k{
              return true
            }else{
              obj[nums[i]]=i
          }            
           
        }else{
            obj[nums[i]]=i
        }
    }

    return false
}


func main()  {
// arr:=[]int{1 ,2, 8, 6 ,7 ,6 }
// largestitem:=findlargestNumber(5,arr)
// fmt.Println(largestitem," largestitem")
// sL, sm:=findSecLargestAndSecondSmallest(5,arr)
// fmt.Printf("second largest %v and second smallest %v",sL, sm)
// 	issorted:=checkifArrayIssorted(5,arr)
// 	fmt.Print(issorted," sorted result")
// 	nums := []int{1, 2, 2, 1}
// 	issortandrot:=isArrySortedAndRotated(4,nums)
// 	fmt.Println(issortandrot,"Is sorted and rotated")
// 	dpres:=removeDuplicates(nums)
// 	fmt.Println(dpres,"removed duplicate")
 	// nums:=[]int{4,1,2,1,2}
	// res:=singleNumber(nums)
	// fmt.Println(res," \nsingle number")
	// moveZeroes(nums)
// var a= []int{1, 2, 3, 4, 6}
// var b =[]int {2, 3, 5}

// 	// unionres:=uniosOfsorted(a,b)
// 	// fmt.Println(unionres)
  
// 	printsubarray(a)

// var a = []int{1, 2, 3, 1, 1, 1, 1}
// var k=3
// 	longesLen:=longestSubarrayWithSumK(a,k )

//     fmt.Println(longesLen)
// nums:=[]int{1,2,3,4,5,6,7}
// k:=3
// rotateLeft(nums,k)
// rotateRight(nums,k)
// OptimizerotateRight(nums,k)

// nums:=[]int{2,0,2,1,1,0}
// sortColors(nums)


// nums:=[]int{2,2,1,1,1,2,2}
// majorityElement(nums)

// nums:=[]int{-1}
// sum:=maxSubArray(nums)
// fmt.Print(sum, "max returns")
// prices:=[]int{7,1,5,3,6,4}
// profit:=maxProfit(prices)
// fmt.Print(profit," max profit")


// nums:=[]int{-1,1}
// rearrangeArray((nums))



// [1,4,7],
// [2,5,8],
// [3,6,9]

//[1,2,3,4,8,12,11,10,9,5,6,7]
// spiralOrder(mat)

// mat:= [][]int{{5,1,9,11},{2,4,8,10},{13,3,6,7},{15,14,12,16}}

// 5  1  9  11
// 2  4  8  10
// 13 3  6   7
// 15 14 12 16

  //  5  2  13  15
  //  1  4   3  14
  //  9  8   6   12
  //  11  10  7 16
// rotate90(mat) 
//  nums := []int{1,1,1}
//  k := 2
// subarraySum(nums,k)

// mat :=[][]int{{1,1,1},{1,0,1},{1,1,1}}
// setZeroes(mat)
// setZeroes2(mat)
// nums1:=[]int{1,1,1,1,1,1,1,1,1,1,4,4}
// nums2:=[]int{1,3,4,4,4,4,4,4,4,4,4}
//  findMedianSortedArrays(nums1,nums2)
// nums:=[]int{1,-1,0}
// subarraySum2(nums, 0)
nums:=[]int{3,4,-1,1}
firstMissingPositive(nums)
}

