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
	n := 10
	fibVal := findFibonacii(n)
	fmt.Println("Fibonacii value of %d is ", n, fibVal)
}

func findFibonacii(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if _, ok := memoTable[n]; ok {
		return memoTable[n]
	}
	fibVal := findFibonacii(n-1) + findFibonacii(n-2)
	memoTable[n] = fibVal
	return memoTable[n]

}
