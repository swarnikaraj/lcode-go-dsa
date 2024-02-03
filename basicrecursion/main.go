package main

import "fmt"

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

}