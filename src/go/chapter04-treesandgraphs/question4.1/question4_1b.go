// Author: Dinesh Appavoo (dineshappavoo) <dinesha.cit@gmail.com>

/*
Algorithm REMOVER_DUPLICATES(list): [Using temp buffer]
 1.Get the input linked list
 2.Create a map of [string]boolean
 4.For i from 1 to len(list) iterate through the nodes nd of the list
		if nd present in the map then remove current node
		else put nd to the map --> map[nd]=true
 5.return list
*/
package main

import (
	"fmt"
	"go/chapter04-treesandgraphs/tree"
	"math"
)

func main() {

	t1 := tree.New(100, 1)
	height := getHeight(t1)
	fmt.Println(height)

	t2 := tree.New(100, 1)
	height = getHeight(t2)
	fmt.Println(height)
}

func getHeight(t *tree.Tree) float64 {
	if t == nil {
		return 0
	}
	return math.Max(getHeight(t.Left), getHeight(t.Right)) + 1
}
