// Author: Dinesh Appavoo (dineshappavoo) <dinesha.cit@gmail.com>

/*
Algorithm KTH_FROM_LAST(list,k):
 1.Get the input linked list and k
 2.if list.Size()<k return nil
 2.for i from 1 to k 
 		Iterate the list and the get the element of k th position kElem
 3.pt=l.Front()
 4.for kElem!=nil;kElem=kElem.Next(),pt=pt.Next()
 		if kElem==nill
 			return pt
 5.return nil
*/
package main
import (
	"container/list"
	"fmt"
)
var sMap map[int]bool
func main() {
	l := list.New()
	l.PushFront(4)
	l.PushFront(5)
	l.PushFront(7)
	l.PushFront(6)
	l.PushFront(5)
	l.PushFront(4)
	l.PushFront(5)
	l.PushFront(7)
	l.PushBack(9)
	kFromLastElem := findKFromLast(l,3)
	fmt.Println(kFromLastElem.Value)
}
func findKFromLast(l *list.List, k int) *list.Element {
	size:=l.Len()
	if size<k {
	return nil
	}
	var elem *list.Element
	elem = l.Front()
	for i:=0;i<k;i++ {
		elem = elem.Next()
	}
	var first *list.Element
	for first = l.Front() ; first!=nil && elem!=nil; elem,first= elem.Next(),first.Next() {
		if elem.Next() == nil {
			return first
		}
	}
	return nil
}
