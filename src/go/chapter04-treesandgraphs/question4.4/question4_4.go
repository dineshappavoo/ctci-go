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
	nodeList = getLevelbasedList(nodeList, t1, 0)
	fmt.Println("")
}

func getLevelbasedList(nodeList []*list.List, t *binarytree.Tree, level int) []*list.List {

	if t==nil {
		return nil
	}
	if nodeList[level]==nil {
			l := list.New()
			l.PushFront(t.Value)
			nodeList[level]=l			
	}else{
		l:=nodeList[level]
		l.PushFront(t.Value)
	}

	adjacents := t.adjacents
	for _,node := range adjacents {
		getLevelbasedList(nodeList, node, level+1)
	}
	return nodeList
}
