// Author: Dinesh Appavoo (dineshappavoo) <dinesha.cit@gmail.com>

/*
 Algorithm DELETE_NODE(node):
 1.Get the pointer to the node
 2.if node==nil || node.next == nil then return false
 3.node.Value = node.next.value
 4.node.next=node.next.next
 5.return true
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()
	l.PushFront(4)
	l.PushFront(5)
	e4 := l.PushFront(7)
	l.PushFront(6)
	l.PushBack(9)
	res := deleteNode(e4)
	fmt.Println(res)

	for e:=l.Front(); e!=nil;e=e.Next() {
		fmt.Println(e.Value)
	}

}

func deleteNode(node *list.Element) bool {

	if node == nil || node.Next() == nil {
		return false
	}

	var next *list.Element 
	next = node.Next()
	node.Value = node.Next().Value
	next = next.Next()
	return true
}
