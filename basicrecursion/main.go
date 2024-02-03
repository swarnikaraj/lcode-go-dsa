package main

import (
	"fmt"
	"strings"
)

// Problem statement
// You are given an integer ‘n’.

// Your task is to return an array containing integers from 1 to ‘n’ (in increasing order) without using loops.

// Example:
// Input: ‘n’ = 5

// Output: 1 2 3 4 5

// Explanation: An array containing integers from ‘1’ to ‘n’ is [1, 2, 3, 4, 5].


func printNNumber(n int) []int{

	

	if n==0{
		return []int{}
	}

	oldarray:=printNNumber(n-1)
	oldarray=append(oldarray, n)
	return oldarray
}




// Problem statement
// You are given an integer ‘n’.
// Your task is to return an array containing integers from ‘n’ to ‘1’ (in decreasing order) without using loops.
// Note:
// In the output, you will see the array returned by you.
// Example:
// Input: ‘n’ = 5

// Output: 5 4 3 2 1

// Explanation: An array containing integers from ‘n’ to ‘1’ is [5, 4, 3, 2, 1].


func printNReverse(n int) []int{
	if n==1{
		return []int{n}
	}

	oldarray:=printNReverse(n-1)
	newarray:=append([]int{n},oldarray...)
	
	return newarray
}



// Problem statement
// You are given an integer ‘n’.

// Print “Golang ” ‘n’ times, without using a loop.

// Example:
// Input: ‘n’  = 4

// Output:
// Coding Ninjas Coding Ninjas Coding Ninjas Coding Ninjas

// Explanation: “Coding Ninjas” is printed 4 times.





func printntimes(n int){

	if n==0{
		return
	}
	fmt.Println("Namaste Golang")
	
	printntimes(n-1)
	
}


// Problem statement
// You are given an integer ‘n’.
// Your task is determining the sum of the first ‘n’ natural numbers and returning it.
// Example:
// Input: ‘n’ = 3

// Output: 6

// Explanation: The sum of the first 3 natural numbers is 1 + 2 + 3, equal to 6.


func printSumOfN(n int) int{

	if n==0{
		return 0
	}
  prevsum:=printSumOfN(n-1) 
   return prevsum + n
}

// Problem statement
// You are given an integer ’n’.
// Your task is to return a sorted array (in increasing order) containing all the factorial numbers which are less than or equal to ‘n’.
// The factorial number is a factorial of a positive integer, like 24 is a factorial number, as it is a factorial of 4.
// Note:
// In the output, you will see the array returned by you.
// Example:
// Input: ‘n’ = 7

// Output: 1 2 6

// Explanation: Factorial numbers less than or equal to ‘7’ are ‘1’, ‘2’, and ‘6’.

func factorialNum(n int) []int{
   
	if n==0 || n==1{
		arr:=[]int{1}
		return arr
	}

	oldarr:=factorialNum(n-1)
	laststorednum:=oldarr[len(oldarr)-1]
    currentfactorial:=laststorednum*n
	oldarr=append(oldarr,currentfactorial )
	return oldarr
}


// Problem statement
// Given an array 'arr' of size 'n'.
// Return an array with all the elements placed in reverse order.
// Note:
// You don’t need to print anything. Just implement the given function.
// Example:
// Input: n = 6, arr = [5, 7, 8, 1, 6, 3]

// Output: [3, 6, 1, 8, 7, 5]

// Explanation: After reversing the array, it looks like this [3, 6, 1, 8, 7, 5].

func reverser(start int, end int, arr []int) []int{
   if start>=end{
	return arr
   }
   arr[start],arr[end]=arr[end],arr[start]
   return reverser(start+1,end-1,arr)
}
func reverseArray(n int, arr []int)  []int{

result :=reverser(0,n-1,arr)
return result
}


// Problem statement
// Determine if a given string ‘S’ is a palindrome using recursion. Return a Boolean value of true if it is a palindrome and false if it is not.

// Note: You are not required to print anything, just implement the function. Example:
// Input: s = "racecar"
// Output: true
// Explanation: "racecar" is a palindrome.

func helperPalindromeChecker(start int, end int, str string) bool{
	if start>=end {
		return true
	}
	if str[start]!=str[end]{return false}
   return helperPalindromeChecker(start+1,end-1,str)
}
func recursivePalindrome(str string) bool{
    checkpalindrom:=helperPalindromeChecker(0,len(str)-1,str)
	return checkpalindrom
}


// Problem statement
// Given an integer ‘n’, return first n Fibonacci numbers using a generator function.

// Example:
// Input: ‘n’ = 5

// Output: 0 1 1 2 3

// Explanation: First 5 Fibonacci numbers are: 0, 1, 1, 2, and 3.
// Note:
// You don't need to print anything. Just implement the given function.


func fibonacciHelper(n int) int{
	if n==0{
		return 0}
   if n==1{
	
	return 0
   }
   if n==2{
	
	return 1
   }

   currentNum:=fibonacciHelper(n-1) + fibonacciHelper(n-2)
   
   
   return currentNum
	
}

func getAllfibonacci(n int){
	 arr:=make([]string,n)
	for i:=1;i<n;i++{
		num:=fibonacciHelper(i)
		arr[i-1]=fmt.Sprint(num)
	}

 fmt.Print(strings.Join(arr," "))


 
}
func main() {
nNumberArrayRes:=printNNumber(6)
fmt.Println(nNumberArrayRes)
printntimes(5)

nReverseResponse:=printNReverse(5)
fmt.Println(nReverseResponse," reverse of n number in array")

sumOfNresponse:=printSumOfN(6)
fmt.Println(sumOfNresponse,"sum of n response")


factorialResponse:=factorialNum(3)
fmt.Println(factorialResponse," factorial of n response")

reversearraResponse:=reverseArray(6,[]int{1,2,3,4,5,6})
fmt.Println(reversearraResponse)


palindromResponse:=recursivePalindrome("annna")
fmt.Println(palindromResponse," palindrome checker response")

getAllfibonacci(4)
}