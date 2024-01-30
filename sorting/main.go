package main

import "fmt"


var myslice = []int{3, 8, 23, 73, 6, 2, 4, 1}

func selectionSort(arr []int) {
	
	for i:=0;i<len(arr);i++{
		selected:=arr[i]
		for j:=i+1;j<len(arr);j++{
          if(arr[j]<=selected){
			arr[i],arr[j]=arr[j] ,arr[i]
		  }
		}
	}

	fmt.Print(arr,"sorted arrray")
	
}

func main() {
selectionSort(myslice)
}
