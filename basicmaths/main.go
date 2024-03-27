package main

import (
	"fmt"
	"math"
	"strconv"
)

func countDigits1(n int) int {
	var count int
	for n > 0 {
		//  digit := n % 10
		count++
		n = n / 10
		fmt.Println(n)
	}
	fmt.Println(count,"in a digit")
return count
	
}
func countDigits2(n int) int {
	strnum:=strconv.Itoa(n)
	 fmt.Println(len(strnum))
return len(strnum)
	
}

func reversedigit(n int) int{
	var reversed int
  for n>0 {
	lastdigit:=n%10
	n=n/10
     reversed=reversed*10 +lastdigit
  }
fmt.Printf("Reverse number %d",reversed)
  return reversed

}

func checkPalindrome(n int) bool{
	fmt.Println(n ==reversedigit(n))
   
	return n ==reversedigit(n)
}

func isArmstrongNumber(n int) bool{
	var sum =0
	var copyn=n
	for n>0{
		lastdidgit:=n%10
      sum=sum+lastdidgit*lastdidgit*lastdidgit

		n=n/10
	}
   fmt.Println(copyn==sum,"Is armstrong")
	return copyn==sum
}

func printAllDivisor1(n int){
	for i:=1;i<=n;i++{
		if n%i==0{
			fmt.Println(i)
		}
	}
}
func printAllDivisor2(n int){
	for i:=1;i<=int(math.Sqrt(float64(n)));i++{
		if(n%i==0 ){
			fmt.Println(i)

			if n/i!=i{
              fmt.Println(n/i)
			}
			
		}
	
	}
}


func isPrime1(n int) bool{
   if n<2 {
	fmt.Println("Not a prime")
    return false 
   }
   for i:=2;i*i<n;i++{
	if n%i==0{
		fmt.Println("Not a prime")
		return false
	}
   }
fmt.Println("A prime")
   return true
}


func printGCD(n1 int, n2 int) int{
    gcd:=1
	for i:=1;i<=int(math.Min(float64(n1), float64(n2)));i++{
		if n1%i==0 && n2%i==0{
			gcd=i
		}
	}
	fmt.Println(gcd,"greatest com factor")
	return gcd
}
// gcd(a,b)=gcd(a%b,b)
// time complexity log base fi of min of n1,n2
func euclidianAlgo(n1 int, n2 int){
     
	for n1>0 && n2>0{
		if(n1>n2){
			n1=n1%n2
		}else{
			n2=n2%n1
		}
	}

	if(n1==0){fmt.Print(n2)
	}else{
		fmt.Print(n1)
	}
}

// reverse a number without converting to a string



func main() {
countDigits1(326978967)
countDigits2(326978967)
reversedigit(123143434)
checkPalindrome(122)
isArmstrongNumber(153)
printAllDivisor1(36)
printAllDivisor2(36)
isPrime1(14)
printGCD(11,22)
euclidianAlgo(11,22)
}

