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
	"go/chapter02-linkedlists/list"
	"go/chapter04-treesandgraphs/binarytree"
)

func main() {
	inArr := []int{4, 5, 7, 8, 9}
	t1 := binarytree.NewMinimalHeightBST(inArr, 0, len(inArr)-1)
	binarytree.InOrderTraverse(t1)
	var nodeList []*list.List
	getLevelbasedList(&nodeList, t1, 0)
	fmt.Println()
	for _, value := range nodeList {
		fmt.Print("[ ")
		for x := value.Front(); x != nil; x = x.Next() {
			fmt.Print(x.Value.(int), " ")
		}
		fmt.Println("]")
	}
}

//Here *[]*list.List used in the function argument to input pass by referrence slice
func getLevelbasedList(nodeList *[]*list.List, t *binarytree.Tree, level int) {
	if t == nil {
		return
	}
	l := list.New()
	if len(*nodeList) == level {
		l = list.New()
		l.PushFront(t.Value)
		*nodeList = append(*nodeList, l)

	} else {
		l = (*nodeList)[level] //index operator doesn't automatically dereference pointers. we need to use parentheses to specify what is dereferenced.
		l.PushFront(t.Value)
	}
	getLevelbasedList(nodeList, t.Left, level+1)
	getLevelbasedList(nodeList, t.Right, level+1)
}
