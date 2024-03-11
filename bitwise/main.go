package main

import (
	"fmt"
)

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

func addBinary(a string, b string) string {
     result := ""
     carry := 0
    i, j := len(a)-1, len(b)-1
   
   for i>=0 || j>=0 || carry>0{
       sum:=carry

    if i >=0 {
         sum=sum + int(a[i]-'0')
          i--
    }
   

    if j>=0{
        sum=sum + int(b[j]-'0')
       j--
    }
     
     fmt.Println(sum," ",sum%2,"sum kya h")
     result=string(sum%2+'0') + result
     carry=sum/2


   }

   return result

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
// more optimized

// You can use the fact that the number of 1's in the binary representation of a number x is equal to the number of 1's in the binary representation of x/2 plus 1 if x is odd, or just equal to the number of 1's in the binary representation of x/2 if x is even.

func countBits2(n int) []int {
    ans := make([]int, n+1)
    
    for i := 1; i <= n; i++ {
        ans[i] = ans[i>>1] + i&1
    }
    
    return ans
}

func main(){
	nums:=[]int{1,2,2,3,3,4,4,5,5}
singleNumber(nums)
countBits(9)
fmt.Println(1+1)
}