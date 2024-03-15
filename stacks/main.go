package main

import "fmt"

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
func (st *Stack) removeItem() bool {
	if st.ptr < 0 {
		return false
	}
	st.ptr--
	return true

}
func (st *Stack) peekItem() int {

	return st.CustomStack[st.ptr]
}

func main() {

	newstack := &Stack{size: 5, CustomStack: make([]int, 5), ptr: -1}
	newstack.addItem(1)
	newstack.addItem(2)
	newstack.addItem(3)

	newstack.removeItem()
	fmt.Println(newstack.peekItem())
	fmt.Println(newstack.CustomStack)

}