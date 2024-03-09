package main

import "fmt"


var myslice = []int{-1, -1, 23, -1, 6, 0, 1, 1}

var myslice2 = []int{1, 2, 3, 4, 5, 6, 7, 8}
var myslice3= []int{4, 2, 5, 1, 3}
func selectionSort(arr []int) {
	n:=len(arr)
	for i:=0;i<n-1;i++{
		minindex:=i
		for j:=i+1;j<n;j++{
          if(arr[j]<arr[minindex]){
			minindex=j
			
		  }
		}
       
		arr[minindex],arr[i]=arr[i] ,arr[minindex]
		
	}

	fmt.Print(arr," selectionSort sorted arrray")
	
}


func bubbleSort(arr[]int){
	n:=len(arr)
	for i:=0;i<n;i++{

		for j:=1;j<n-i;j++{
             if(arr[j]<arr[j-1]){
				arr[j],arr[j-1]=arr[j-1],arr[j]
			 }
		}
	}
    fmt.Print(arr," bubbleSort sorted arrray")

}


func insertionsort(arr []int){

n:=len(arr)
for i:=0;i<=n-2;i++{


	for j:=i+1; j>0; j--{

		if(arr[j]<arr[j-1]){
          arr[j], arr[j-1]=arr[j-1], arr[j]
		}else{
			break
		}
       
	}


}

 
 fmt.Print(arr,"insertion sort")
}


func cyclicSort(arr []int){
	i:=0
	for i<len(arr){
       value:=arr[i]
		index:=value-1

		if(i==arr[index]){
			i++
		}else{
           arr[i], arr[index]=arr[index],arr[i]
		}
	}
 fmt.Print(arr,"cyclic sort")
}	


func main() {
// selectionSort(myslice)
// bubbleSort(myslice2)
insertionsort(myslice)
cyclicSort(myslice3)
}
