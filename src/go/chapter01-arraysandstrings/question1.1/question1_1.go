// Author: Dinesh Appavoo (dineshappavoo) <dinesha.cit@gmail.com>

/*
Algorithm IS_DUPLICATE_EXIST(str):
 1.Get the input string
 2.Create a map of [string]boolean
 4.For i 1 to n iterate through the characters ch of the string
		if ch present in the map then return false
		else put ch to the map --> map[ch]=true
	5.return true
*/
package main

import (
	"fmt"
)

var sMap map[string]bool

func main() {
	str := "Helo World!"
	fmt.Println(str)
	fmt.Println(is_duplicate_present(str))
}

func is_duplicate_present(str string) bool {

	sMap = make(map[string]bool)
	var ch string
	for i := 0; i < len(str); i++ {
		ch = string([]rune(str)[i])
		if sMap[ch] {
			return false
		} else {
			sMap[ch] = true
		}
	}
	return true
}
