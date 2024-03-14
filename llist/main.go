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

func (ll *LinkedList) AppendData(data int){
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

func (ll *LinkedList) AppendAtFirst(data int){
	newnode:= &Node{data:data}

	if ll.head==nil{
		ll.head=newnode

	}else{
		newnode.next=ll.head
		ll.head=newnode
	}
	if ll.tail==nil{
		ll.tail=newnode
	}
	ll.size++
}


func (ll *LinkedList) printll(){
	curr:=ll.head
	for curr.next!=nil{
		fmt.Printf("node data= %v node next = %v",curr.data, curr.next)
		curr=curr.next
	}
	fmt.Println()

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     data int
 *     next *ListNode
 * }
 */
func removeElements(head *Node, data int) *Node {
    for head != nil && head.data == data {
        head = head.next
    }
    
    if head == nil {
        return nil
    }
    
    prev := head
    curr := head.next
    
    for curr != nil {
        if curr.data == data {
            prev.next = curr.next
        } else {
            prev = curr
        }
        curr = curr.next
    }
    
    return head
}

func (ll *LinkedList)addElementAt(index int, data int){
	
	if index==0{
       ll.AppendAtFirst(data)
	   return
	}else if index==ll.size{
		ll.AppendData(data)
	}else{
     
     temp:=ll.head
	 for i:=1;i<index;i++{
     temp=temp.next
	 }
     newnode:=&Node{data:data, next: temp.next}
	  temp.next=newnode
	}
    


	
}


func main() {
	ll := LinkedList{}
   ll.AppendData(1)
    ll.AppendData(2)
	ll.AppendData(3)
	ll.AppendAtFirst(5)
	ll.addElementAt(3,86)
//  fmt.Println((ll))
	ll.printll()
}