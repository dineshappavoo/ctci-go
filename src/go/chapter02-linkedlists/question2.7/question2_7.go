// Author: Dinesh Appavoo (dineshappavoo) <dinesha.cit@gmail.com>

/*
 Algorithm IS_PALINDROME(list):
 1.Get the input list
 2.if list==nil return nil
 3.Declare a stack s
 4.Use slow and fast runner technique to add the elements to the list till the half of the list
 5.
 3.Make two pointers runner and current
 4.move current once and runner twice on every iteration
 5.
 3.add elements and maintain carry from lists until both the lists has elements
 4.Add carry to the continuing list and add value to result list if the lists are of different size
 5.return the result list
*/

package main

import (
	"container/list"
	"fmt"
)

var head *list.Element

func main() {

	l := list.New()
	l.PushFront(4)
	l.PushFront(5)
	l.PushFront(4)
	l.PushFront(5)
	l.PushFront(4)

	m := list.New()
	m.PushFront(3)
	m.PushFront(2)
	m.PushFront(8)
	m.PushFront(6)
	res := isPalindrome(l)
	stackImplementation()
	fmt.Println(res)
	res = isPalindromeUsingStack(l)
	fmt.Println(res)
}

//Reverse list function
func reverseList(l *list.List) *list.List {
	m := list.New()

	for e := l.Front(); e != nil; e = e.Next() {
		m.PushFront(e.Value.(int))
	}
	return m
}
/*
func reverseList1(current *list.Element) {
	if (current == nil) {
		return
	}

	if (current.Next()==nil) {
		head = current
		return
	}
	reverseList1(current.Next())
	current.Next().Next() = current
	current.Next() = nil
}*/

//Palinfrome function using reverse list approach - Iterative
func isPalindrome(l *list.List) bool {

	rev := reverseList(l)
	for e, f := l.Front(), rev.Front(); e != nil && f != nil; e, f = e.Next(), f.Next() {
		if e.Value.(int) != f.Value.(int) {
			return false
		}
	}
	return true
}


//Stack implementation
type Stack struct {
	top  *Node
	size int
}

type Node struct {
	value interface{}
	next  *Node
}

func (s *Stack) Length() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) Push(val interface{}) {
	s.top = &Node{val, s.top}
	s.size++
}

func (s *Stack) Peek() interface{} {
	return s.top.value
}

func (s *Stack) Pop() (val interface{}) {
	if s.size > 0 {
		val, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return ""
}

//Test Stack operations
func stackImplementation() {
	st := new(Stack)

	st.Push("1")
	st.Push("2")
	st.Push("3")
	st.Push("4")
	st.Push("5")
	st.Push("6")

	for st.Length() > 0 {
		fmt.Printf("%s ", st.Pop().(string))
	}
	fmt.Println()
}

func isPalindromeUsingStack(l *list.List) bool {

	if l == nil {
		return false
	}
	for e := l.Front(); e != nil; e = e.Next() {
		n := e.Value.(int)
		fmt.Println(n)
	}
	st := new(Stack)
	slow := l.Front()
	fast := l.Front()

	for ; fast != nil || fast.Next() != nil ; {
		fmt.Println("test")
		st.Push(slow.Value.(int))
		slow = slow.Next()
		fmt.Println("Fastets")
		fast = fast.Next().Next()
	}

	fmt.Println("Testing 1.1")

	//To ignore the middle element in case of odd length
	if fast.Next() == nil {
		slow = slow.Next()
	}

	fmt.Println("Testing 1")
	for ; slow != nil; {
		m := st.Pop().(int)
		n := slow.Value.(int)
		if (m!=n) {
			return false
		}
	}
	return true	
}
