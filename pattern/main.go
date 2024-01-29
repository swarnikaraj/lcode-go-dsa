package main

import (
	"fmt"
)

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


func pattern11(n int){
    var startWith int
	for i:=1;i<=n;i++{
		cols:=i
	   if isOdd:=i%2==1 ; isOdd{
        startWith=1
	   }else{
		startWith=0
	   }
        printVal:=startWith
		for j:=1;j<=cols;j++{
           fmt.Print(printVal)
		 printVal=  printVal ^ 1
		}

		fmt.Print("\n")
	}
}


// Problem statement
// Aryan and his friends are very fond of the pattern. They want to make the Reverse N-Number Crown for a given integer' N'.

// Given 'N', print the corresponding pattern.

// Example:
// Input: ‘N’ = 3

// Output: 

// 1         1
// 1 2     2 1
// 1 2 3 3 2 1

func pattern12(n int){
	for i:=1;i<=n;i++{

		cols:=i
		spaces:=2*n-2*i
		

		for j:=1;j<=cols;j++{
         fmt.Print(j)
		}
		for k:=1;k<=spaces;k++{
			fmt.Print(" ")
		}
		for l:=cols;l>=1;l--{
			fmt.Print(l)
		}

		fmt.Print("\n")
	}
}








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


func pattern16(n int){
	 printVal := 'A'
	for i:=1;i<=n;i++{
       
	  cols:=i
         for j:=1;j<=cols;j++{
			fmt.Print(string(printVal))
		 }
        printVal++
		
		 fmt.Print("\n")
         
	   
	 

	}
}



// Problem statement
// Sam is curious about Alpha-Hills, so he decided to create Alpha-Hills of different sizes.

// An Alpha-hill is represented by a triangle, where alphabets are filled in palindromic order.

// For every value of ‘N’, help sam to return the corresponding Alpha-Hill.

// Example:
// Input: ‘N’ = 3

// Output: 
//     A
//   A B A
// A B C B A

 func pattern17(n int){
	
	for i:=1;i<=n;i++{
       rows:=2*i-1
        printVal:='A'
         
		   median:=(rows+1)/2
		   spaces:=n-median
       for j:=1;j<=spaces;j++{
         fmt.Print(" ")
        }
      for j:=1;j<=rows;j++{
        fmt.Print(string(printVal))
         if(j<median){
           printVal++
		 }else{
			printVal--
		 }
		
	  }

	  fmt.Print("\n")
	}
 }




// Problem statement
// Sam is researching on Alpha-Triangles. So, he needs to create them for different integers ‘N’.

// An Alpha-Triangle is represented by the triangular pattern of alphabets in reverse order.

// For every value of ‘N’, help sam to print the corresponding Alpha-Triangle.

// Example:
// Input: ‘N’ = 3

// Output: 
// C
// C B 
// C B A


func pattern18(n int){
	  
	alphas:="ABCDEFGFIJKLMNOPQRSTUVWXYZ"
	for i:=1;i<=n;i++{
		 alphindex:=n-1
		 

		 rows:=i

		 for j:=1;j<=rows;j++{
			fmt.Print(string(alphas[alphindex]))
			alphindex=alphindex-1
		 }
		 fmt.Print("\n")




	}
}


// Problem statement
// Sam is curious about symmetric patterns, so he decided to create one.

// For every value of ‘N’, return the symmetry as shown in the example.

// Example:
// Input: ‘N’ = 3

// Output: 
// * * * * * * 
// * *     * * 
// *         * 
// *         * 
// * *     * * 
// * * * * * * 



func pattern19(n int){
    // upper paer
	 colsU:=n
	for u:=1;u<=n;u++{
     //left start
    
	 space:=2*n-2*colsU
	 for left:=1;left<=colsU;left++{
     fmt.Print("*")
	 }
     //spaces
     for spc1:=1;spc1<=space;spc1++{
		fmt.Print(" ")
	 }
	 //right stars
	  for left:=1;left<=colsU;left++{
	  fmt.Print("*")	
	 }
	 colsU=colsU-1
     fmt.Print("\n") 
	}

    colsB:=1

    for l:=1;l<=n;l++{

		//left 
   spaces2:=2*n-2*colsB
		for l:=1;l<=colsB;l++{
			fmt.Print("*")
		}
/// spaces
		for spc2:=1;spc2<=spaces2;spc2++{
			fmt.Print(" ")
		}
 

for l:=1;l<=colsB;l++{
			fmt.Print("*")
		}
		colsB++
         fmt.Print("\n")
	}

}

// Problem statement
// Sam is curious about symmetric patterns, so he decided to create one.

// For every value of ‘N’, return the symmetry as shown in the example.

// Example:
// Input: ‘N’ = 3

// Output: 
// *         *
// * *     * *
// * * * * * *
// * *     * *
// *         *

func pattern20(n int){
   
    colsB:=1
    rows:=2*n-1
    for l:=1;l<=rows;l++{

		//left 
      spaces2:=2*n-2*colsB
		for l:=1;l<=colsB;l++{
			fmt.Print("*")
		}
/// spaces
		for spc2:=1;spc2<=spaces2;spc2++{
			fmt.Print(" ")
		}
 

     for l:=1;l<=colsB;l++{
			fmt.Print("*")
		}

if(l<n){
colsB++
}else{
	colsB--
}
		
		
         fmt.Print("\n")
	}

}


// Problem statement
// Ninja has been given a task to print the required star pattern and he asked your help for the same.

// You must return an ‘N’*’N’ matrix corresponding to the given star pattern.

// Example:
// Input: ‘N’ = 4

// Output: 

// ****
// *  *
// *  *
// ****


func pattern21(n int){

	for i:=1;i<=n;i++{
		fmt.Print("*")
	}
   fmt.Print("\n")
	for k:=1;k<=n-2;k++{
   
	fmt.Print("*")
    for j:=1;j<=n-2;j++{
		fmt.Print(" ")
	}
    fmt.Print("*")
	fmt.Print("\n")
	}
   


	for i:=1;i<=n;i++{
		fmt.Print("*")
	}
}


// Problem statement
// Ninja has been given a task to print the required pattern and he asked for your help with the same.

// You must print a matrix corresponding to the given number pattern.

// Example:
// Input: ‘N’ = 4

// Output: 

// 4444444
// 4333334
// 4322234
// 4321234
// 4322234
// 4333334
// 4444444


func pattern22(n int){
	size:=2*n-1
	matt:=make([][]int,size)
    top:=0
	bot:=size-1
	left:=0
	right:=size-1
	starter:=n
    var count int
     for i := 0; i < size; i++ {
        matt[i] = make([]int, size)
     }

	for count<size*size{

		for i:=left;i<=right;i++{
		matt[top][i]=starter
		count++
	
	  }
	  top++

	  for j:=top;j<=bot;j++{
		matt[j][right]=starter
		count++
		
	  }
	  right--
	  for k:=right;k>=left;k--{
		matt[bot][k]=starter
		count++

	  }
	  bot--

	  for l:=top;l<=bot;l++{
		matt[l][left]=starter
		count++
	  }
	  left++

	 starter--
	 

	}
     
	 
   
	
   for row:=0;row<size;row++{
      for col:=0;col<size;col++{
		fmt.Print(matt[row][col])
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
	 pattern11(4)
     fmt.Println("============================")
	 pattern12(4)
	 fmt.Println("============================")
	 pattern13(4)
     fmt.Println("============================")
	 pattern14(4)
     fmt.Println("============================")
	 pattern15(4)
     fmt.Println("============================")
	 pattern16(4)
	 fmt.Println("============================")
     

	 pattern17(4)
	 fmt.Println("============================")
	 pattern18(3)
	 fmt.Println("============================")

	 pattern19(4)

	 fmt.Println("============================")

	 pattern20(4)

	 fmt.Println("============================")
     pattern21(5)
     fmt.Println("============================")
	 pattern22(4)



     
}