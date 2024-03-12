package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (ll *LinkedList) appendData(data int){
      newnode:=&Node{data:data}

      if ll.head==nil{
		ll.head=newnode
	  }else{
		curr:=ll.head
		for curr.next!=nil{
        
		   curr=curr.next
		}
		curr.next=newnode
	  }
    ll.size++
	ll.tail=newnode
}

func (ll *LinkedList) appendAtFirst(data int){
	newnode:= &Node{data:data}

	if ll.head==nil{
		ll.head=newnode

	}else{
		newnode.next=ll.head
		ll.head=newnode
	}
	ll.size++
}


func (ll *LinkedList) printll(){
	curr:=ll.head
	for curr.next!=nil{
		fmt.Printf("node val= %v node next = %v",curr.data, curr.next)
		curr=curr.next
	}
	fmt.Println()

}
func main() {
	ll := LinkedList{}
   ll.appendData(1)
    ll.appendData(2)
	ll.appendData(3)
	ll.appendAtFirst(5)
	// fmt.Println((ll))
	ll.printll()
}