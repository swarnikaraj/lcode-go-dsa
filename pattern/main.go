package main

import "fmt"

// Problem statement
// Sam is making a forest visualizer. An N-dimensional forest is represented by the pattern of size NxN filled with ‘*’.

// For every value of ‘N’, help sam to print the corresponding N-dimensional forest.

// Example:
// Input: ‘N’ = 3

// Output:
// * * *
// * * *
// * * *


func pattern1(n int){
	for i:=1 ;i<=n;i++{
       for j:=1 ; j<=n ;j++{
		fmt.Print("*")
	   }
	   fmt.Print("\n")
	}
}


// Problem statement
// Sam is making a forest visualizer. An N-dimensional forest is represented by the pattern of size NxN filled with ‘*’.

// An N/2-dimensional forest is represented by the lower triangle of the pattern filled with ‘*’.

// For every value of ‘N’, help sam to print the corresponding N/2-dimensional forest.

// Example:
// Input:  ‘N’ = 3

//  
// Output: 
// * 
// * *
// * * *

 func pattern2(n int){
  for i:=1;i<=n;i++{
	for j:=1;j<=i;j++{
		fmt.Print("*")
	}
	fmt.Print("\n")
  }
}

// Problem statement
// Sam is making a Triangular painting for a maths project.

// An N-dimensional Triangle is represented by the lower triangle of the pattern filled with integers starting from 1.

// For every value of ‘N’, help sam to print the corresponding N-dimensional Triangle.

// Example:
// Input: ‘N’ = 3

// Output: 
// 1
// 1 2 
// 1 2 3
// Detailed explanation ( Input/output format, Notes, Images )
// Constraints :
// 1  <= N <= 25
// Time Limit: 1 sec
// Sample Input 1:
// 3
// Sample Output 1:
// 1
// 1 2 
// 1 2 3
// Sample Input 2 :
// 1
// Sample Output 2 :
// 1

func pattern3( n int){
	for i:=1;i<=n;i++{
		for j:=1;j<=i;j++{
			fmt.Print(j," ")
		}
		fmt.Print("\n")
	}
}


// Problem statement
// Sam is making a Triangular painting for a maths project.

// An N-dimensional Triangle is represented by the lower triangle of the pattern filled with integers representing the row number.

// For every value of ‘N’, help sam to print the corresponding Triangle.

// Example:
// Input: ‘N’ = 3
// Output: 
// 1
// 2 2 
// 3 3 3

func pattern4(n int){
	for i:=1;i<=n;i++{
		for j:=1;j<=i;j++{
          fmt.Print(i," ")
		}
		fmt.Print("\n")
	}
}


// Problem statement
// Sam is planting trees on the upper half region (separated by the left diagonal) of the square shared field.
// For every value of ‘N’, print the field if the trees are represented by ‘*’.
// Example:
// Input: ‘N’ = 3

// Output: 
// * * *
// * *
// *


func pattern5(n int){
   for i:=0;i<n;i++{
	for j:=0;j<n-i;j++{
		fmt.Print("* ")
	}
	fmt.Print("\n")
   }
}

// Problem statement
// Aryan and his friends are very fond of the pattern. For a given integer ‘N’, they want to make the Reverse N-Number Triangle.

// Example:
// Input: ‘N’ = 3

// Output: 

// 1 2 3
// 1 2
// 1

func pattern6(n int){
   for i:=0;i<n;i++{
        for j:=0;j<n-i;j++{
			fmt.Print(j+1," ")
		}
		fmt.Print("\n")
   }
}

// Problem statement
// Ninja was very fond of patterns. For a given integer ‘N’, he wants to make the N-Star Triangle.

// Example:
// Input: ‘N’ = 3

// Output: 

//   *
//  ***
// *****
func pattern7(n int){
 for i:=0;i<=n;i++{
	
	strs:=2*i-1

	spc:=((2*n-1)-strs)/2
    for k:=1;k<=spc;k++{
		fmt.Print(" ")
	 }

     for j:=1;j<=strs;j++{
		fmt.Print("*")
	 }

	 for k:=1;k<=spc;k++{
		fmt.Print(" ")
	 }
	 fmt.Print("\n")
 }
}


// Problem statement
// Ninja was very fond of patterns. For a given integer ‘N’, he wants to make the Reverse N-Star Triangle.

// Example:
// Input: ‘N’ = 3

// Output: 

// *****
//  ***
//   *

func pattern8(n int){
	for i:=0;i<n;i++{
         
		stars:=2*n-(2*i+1)
		 spaces:=((2*n-1)-stars)/2
     for j:=1;j<=spaces;j++{
			fmt.Print(" ")
		}
		for j:=1;j<=stars;j++{
			fmt.Print("*")
		}

		for j:=1;j<=spaces;j++{
				fmt.Print(" ")
		}
       fmt.Print("\n")

	}
}

// Problem statement
// Ninja was very fond of patterns. For a given integer ‘N’, he wants to make the N-Star Diamond.

// Example:
// Input: ‘N’ = 3

// Output: 

//   *
//  ***
// *****
// *****
//  ***
//   *


func pattern9(n int){
	// first pyramid
	for i:=0;i<=n;i++{
	
	strs:=2*i-1

	spc:=((2*n-1)-strs)/2
    for k:=1;k<=spc;k++{
		fmt.Print(" ")
	 }

     for j:=1;j<=strs;j++{
		fmt.Print("*")
	 }

	 for k:=1;k<=spc;k++{
		fmt.Print(" ")
	 }
	 fmt.Print("\n")
   }

   // upside down
		for i:=0;i<n;i++{
         
		stars:=2*n-(2*i+1)
		 spaces:=((2*n-1)-stars)/2
        for j:=1;j<=spaces;j++{
			fmt.Print(" ")
		}
		for j:=1;j<=stars;j++{
			fmt.Print("*")
		}

		for j:=1;j<=spaces;j++{
				fmt.Print(" ")
		}
       fmt.Print("\n")

	}
}

// Problem statement
// Ninja was very fond of patterns. For a given integer ‘N’, he wants to make the N-Star Rotated Triangle.

// Example:
// Input: ‘N’ = 3

// Output: 

// *
// **
// ***
// **
// *
func pattern10(n int){

    rows:=2*n-1
	var stars int
	for i:=1;i<=rows;i++{
		
      if(i<=n){
    stars=i
	
	  }else{
    stars =stars-1
	  }

	  for j:=1;j<=stars;j++{
		fmt.Print("*")
	  }
	  fmt.Print("\n")

	} 

}


// Problem statement
// Aryan and his friends are very fond of the pattern. For a given integer ‘N’, they want to make the N-Binary Number Triangle.

// You are required to print the pattern as shown in the examples below.

// Example:
// Input: ‘N’ = 3

// Output: 

// 1
// 0 1
// 1 0 1














// Problem statement
// Aryan and his friends are very fond of patterns. For a given integer ‘N’, they want to make the Increasing Number Triangle.

// Example:
// Input: ‘N’ = 3

// Output: 

// 1
// 2 3
// 4 5 6
func pattern13(n int){

	for i:=1;i<=n;i++{

		printedVal:=1
        
		cols:=i

		for j:=1;j<=cols;j++{
			fmt.Print(printedVal)
			printedVal++
		}
       fmt.Print("\n")
	}


}


// Problem statement
// Aryan and his friends are very fond of patterns. For a given integer ‘N’, they want to make the Increasing Letter Triangle.

// Example:
// Input: ‘N’ = 3

// Output: 

// A
// A B
// A B C


func pattern14(n int){
  

	for i:=1;i<=n;i++{
      
       cols:=i;
	   printval:='A'
	   for j:=1;j<=cols;j++{
		fmt.Print(string(printval))
		printval++
	   }
	   fmt.Print("\n")

	}
}


// Problem statement
// Aryan and his friends are very fond of patterns. For a given integer ‘N’, they want to make the Reverse Letter Triangle.

// You must print a matrix corresponding to the given Reverse Letter Triangle.

// Example:
// Input: ‘N’ = 3

// Output: 

// A B C
// A B
// A


func pattern15(n int){
	for i:=0;i<=n;i++{
       var cols=n-i+1
       printVal:='A'
	
         for j:=1;j<=cols;j++{
			fmt.Print(string(printVal))
			printVal++
		 }
		 fmt.Print("\n")
         
	   
	 

	}
}
func main(){
	pattern1(3)
	fmt.Println("============================")
	pattern2(4)
	fmt.Println("============================")
	pattern3(4)
	fmt.Println("============================")
    pattern4(4)
	fmt.Println("============================")
	pattern5(3)
	fmt.Println("============================")
	pattern6(4)
	fmt.Println("============================")
	pattern7(4)
	fmt.Println("============================")
	pattern8(4)
     fmt.Println("============================")
	 pattern9(3)
     fmt.Println("============================")
	 pattern10(3)
 fmt.Println("============================")
	 pattern13(4)
fmt.Println("============================")
	 pattern14(4)
fmt.Println("============================")
	 pattern15(4)

}