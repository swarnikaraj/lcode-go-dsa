package main

import (
	"fmt"
	"log"
)

type Stack struct {
	size        int
	CustomStack []int
	ptr         int
}

func (st *Stack) addItem(n int) bool {
	if st.size-1 == st.ptr {
		return false
	}
	st.ptr++
	st.CustomStack[st.ptr] = n
	return true
}
func (st *Stack) removeItem() int {
	if st.ptr < 0 {
		log.Fatal("Empty stack")
	}
	st.ptr--
	return st.CustomStack[st.ptr]

}
func (st *Stack) peekItem() int {
	if st.ptr==-1{
		log.Fatal("Empty stack")
	}

	return st.CustomStack[st.ptr]
}

func main() {

	newstack := &Stack{size: 5, CustomStack: make([]int, 5), ptr: -1}
	newstack.addItem(1)
	newstack.addItem(2)
	newstack.addItem(3)
    newstack.addItem(2)
	newstack.addItem(3)
	newstack.addItem(1)
	newstack.addItem(2)
	newstack.removeItem()
	fmt.Println(newstack.peekItem())
	fmt.Println(newstack.CustomStack)
    
}