package main

import (
	"fmt"
	"log"
)

type Queue struct {
	CustomQueue []int
	size        int
	end         int
}

func (q *Queue) addItem(n int) bool {
	if q.size == q.end {
		return false
	}

	q.CustomQueue[q.end] = n
	q.end++
	
	return true
}
func (q *Queue) removeItem() int {
	if q.end == 0 {
		log.Fatal("Empty stack")
	}
	removed := q.CustomQueue[0]

	for i := 1; i < q.size; i++ {
      q.CustomQueue[i-1]=q.CustomQueue[i]
	}
    q.end--
	return removed

}
func (q *Queue) peekItem() int {
	if q.size == 0 {
		log.Fatal("Empty stack")
	}

	return q.CustomQueue[0]                  
}

func main() {
	newque := &Queue{CustomQueue: make([]int, 5), size: 5, end: 0}

	newque.addItem(1)
	newque.addItem(2)
	newque.addItem(3)
	newque.addItem(4)
	newque.addItem(5)


	fmt.Println(newque.size)
	fmt.Println(newque.CustomQueue)
	fmt.Println(newque.removeItem())
	fmt.Println(newque.CustomQueue)
	

}