// Author: Dinesh Appavoo (dineshappavoo) <dinesha.cit@gmail.com>

/*
 Algorithm ADD_LISTS(list1,list2):
 1.Get the input list1 and list2
 2.if list1==nil && list2==nil then return nil
 	else  return list1==nil?list2:list1
 3.add elements and maintain carry from lists until both the lists has elements
 4.Add carry to the continuing list and add to result list if the lists are of different size
 5.return the result list
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
	l.PushFront(7)
	l.PushFront(6)
	l.PushBack(9)
	res := splitAroundX(l, 6)
	for e := res.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

//ERROR

func splitAroundX(l *list.List, x int) *list.List {

	if l == nil {
		return nil
	}
	lThanX := list.New()
	gThanX := list.New()

	for e := l.Front(); e != nil; e = e.Next() {
		m := e.Value.(int)
		if m < x {
			lThanX.PushBack(m)
		} else {
			gThanX.PushBack(m)
		}
	}
	lThanX.PushBackList(gThanX)
	return lThanX
}
