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
	"go/chapter04-treesandgraphs/queue"
	"go/chapter04-treesandgraphs/tree"
	"reflect"
)

func main() {

	t1 := tree.NewTree()
	pathAvailable := isPathAvailable(t1, t1)
	fmt.Println("Is Path Availabe ? ", pathAvailable)
}

func isPathAvailable(t1 *tree.Tree, t2 *tree.Tree) bool {
	if t1 == nil || t2 == nil {
		return false
	}
	q := queue.New()

	if t1.Visited == false {
		q.Push(t1)
		t1.Visited = true
		if t1.Value == t2.Value {
			return true
		}
	}
	for !q.IsEmpty() {
		node := q.Poll()

		//go reflect package used the get the array of inteface from the struct
		refValue := reflect.ValueOf(&node).Elem().FieldByName(string("Adjacents"))
		adjNodes := refValue.Interface().([]*tree.Tree)

		for _, n := range adjNodes {
			if n.Visited == false {
				if n.Value == t2.Value {
					return true
				}
				n.Visited = true
				q.Push(n)
			}
		}
	}
	return false
}
