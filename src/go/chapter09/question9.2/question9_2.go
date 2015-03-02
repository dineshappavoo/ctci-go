// Author: Dinesh Appavoo (dineshappavoo) <dinesha.cit@gmail.com>

/*
Algorithm IS_DUPLICATE_EXIST(str):
 1.Get the input string
 2.Create a map of [string]boolean
 4.For i from 1 to n iterate through the characters ch of the string
		if ch present in the map then return false
		else put ch to the map --> map[ch]=true
	5.return true
*/
package main

import (
	"fmt"
)

var memoTable = make(map[int]int)

func main() {
	n := 35
	noOfWays := findPossibleWaysUsingDP(n)
	fmt.Println("Possible Ways ", noOfWays)
}

//Find possible no of ways without dynamic programming
func findPossibleWays(n int) int {
	//Base case
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	return findPossibleWays(n-1) + findPossibleWays(n-2) + findPossibleWays(n-3)
}

//Find possible no of ways using dynamic programming
func findPossibleWaysUsingDP(n int) int {
	//Base case
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	if _, ok := memoTable[n]; ok {
		return memoTable[n]
	}
	ways := findPossibleWaysUsingDP(n-1) + findPossibleWaysUsingDP(n-2) + findPossibleWaysUsingDP(n-3)
	memoTable[n] = ways
	return memoTable[n]
}
