package main

import "fmt"


var myslice = []int{-1, -1, 23, -1, 6, 0, 1, 1}

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
	for i:=0;i<n-1;i++{

		for j:=i+1;j<n;j++{
             if(arr[j]<arr[i]){
				arr[j],arr[i]=arr[i],arr[j]
			 }
		}
	}
    fmt.Print(arr," bubbleSort sorted arrray")

}

func main() {
selectionSort(myslice)
bubbleSort(myslice)
}
