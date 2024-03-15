package main

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

}