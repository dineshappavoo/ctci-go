// Author: Dinesh Appavoo (dineshappavoo) <dinesha.cit@gmail.com>

/*Algorithm BST_MINIMAL_HEIGHT(arr, low, high): 
1.Get the input array arr
2.find the mid element using low and high . mid =(low+high)/2
3.t.Value = arr[mid]
4.t.Left=BST_MINIMAL_HEIGHT(arr, low, mid)
5.t.Right=BST_MINIMAL_HEIGHT(arr, mid+1, high)
6.return t
*/
package main

import (
	"fmt"
	"go/chapter04-treesandgraphs/binarytree"
	"go/chapter02-linkedlists/list"
)

func main() {

	inArr := []int{4, 5, 7, 8, 9}
	t1 := getMinimalBST(inArr, 0, len(inArr)-1)
	binarytree.InOrderTraverse(t1)
	nodeList := []*list.List
	getLevelbasedList(nodeList, t1, 0)

	for _,value := range nodeList {
		for _ lVal := range value {
			fmt.Print(lVal)
		}
		fmt.Println()
	}
}

func getLevelbasedList(nodeList []*list.List, root *binarytree.Tree, level int) []*list.List{

	if t==nil {
		return
	}
	var nodeList []*list.List
	var parents []int

	parents = append(parents, root.Value.(int))

	var current []int

	for parents != nil {
		current = append(current[:0], current[1:]...)
	}
	 := list.New()
	if len(nodeList)==level {
			l = list.New()
			l.PushFront(t.Value)
	}else{
		l=nodeList[level]
		l.PushFront(t.Value)
	}
	nodeList[level]=l
	getLevelbasedList(nodeList, t.Left, level+1)
	getLevelbasedList(nodeList, t.Right, level+1)
}
