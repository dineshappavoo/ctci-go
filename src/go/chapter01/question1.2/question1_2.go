// Author: Dinesh Appavoo (dineshappavoo) <dinesha.cit@gmail.com>

/*
 Algorithm REVERSE_STRING(S):
 1.Get the input string
 2.Find the length of the string
 3.Create a map of [string]boolean
 4.For i 1 to n iterate through the characters ch of the string
		if ch present in the map then return false
		else put ch to the map --> map[ch]=true
	5.return true
*/
package main

import (
	"fmt"
)

func main() {
	str := "Helo World!"
	fmt.Println("Input String :", str)
	fmt.Println("Reversed String ", reverseString(str))
}

func reverseString(s string) string {

	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)

}
