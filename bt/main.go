package main

import "fmt"



func addtoArray(arr *[]int , i int){
  if i==5{
    return
  }
  fmt.Println((*arr)[len(*arr)-1])
  *arr=append(*arr,i)
  addtoArray(arr,i+1)
}
func addToArray2(arr []int, i int) {
    if i == 5 {
        return
    }
    fmt.Println(arr[len(arr)-1])
    arr = append(arr, i)
    addToArray2(arr, i+1)
}




func main() {
	arr:=[]int{}
    fmt.Println(arr)
   addToArray2([]int{0}, 1)
   fmt.Println(arr,"arr")
}
