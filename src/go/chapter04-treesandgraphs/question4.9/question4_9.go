// Author: Dinesh Appavoo (dineshappavoo) <dinesha.cit@gmail.com>

/*Algorithm IS_T2_PRESENT_IN_T1(t1, t2): 
1.Get the input tree t1 and t2
2.if t1==nil && t2 == nil then return false
3.if t1 == nil || t2 == nil  then return false
4.if t1.Value == t2.Value the return IS_ALL_NODES_AVAILABLE_IN-T1(t1, t2)
5.return  IS_T2_PRESENT_IN_T1(t1.Left, t2.Left) ||IS_T2_PRESENT_IN_T1(t1.Right, t2.Right)

Algorithm IS_ALL_NODES_AVAILABLE_IN-T1(t1, t2):
1.1.Get the input tree t1 and t2
2.if t1==nil && t2 == nil then return true
3.if t1 != nil || t2 == nil  then return true
4.if t1.Value != t2.Value then return false
5.return IS_ALL_NODES_AVAILABLE_IN-T1(t1.Left, t2.Left) && IS_ALL_NODES_AVAILABLE_IN-T1(t1.Right, t2.Right)

*/
package main

import (
	"fmt"
	"go/chapter04-treesandgraphs/binarytree"
)

func main() {

	inArr := []int{4, 5, 7, 8, 9, 10, 12, 14, 15, 17, 19, 20}
	t1 := binarytree.NewMinimalHeightBST(inArr, 0, len(inArr)-1)

	findAndPrintPath(t1)
}

func findAndPrintPath(t *binarytree.Tree) {
	height := binarytree.Height(t)
	fmt.Println("Height", height)
	var path = make([]int, int(height))
	findPathEqualsSum(t, path, 15, 0)
}

func findPathEqualsSum(t *binarytree.Tree, path []int, sum int, level int) {
	if t == nil {
		return
	}
	fmt.Println("Level", level)
	path[level] = t.Value
	val := 0
	for i := level; i >= 0; i++ {
		val += path[i]
		if val == sum {
			printPath(path, i, level)
		}
	}
	findPathEqualsSum(t.Left, path, sum, level+1)
	findPathEqualsSum(t.Right, path, sum, level+1)
}

func printPath(path []int, start int, end int) {
	for i := start; i <= end; i++ {
		fmt.Println(path[i])
	}
}
