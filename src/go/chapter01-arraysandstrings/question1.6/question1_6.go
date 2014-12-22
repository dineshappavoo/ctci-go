// Author: Dinesh Appavoo (dineshappavoo) <dinesha.cit@gmail.com>

// Write a method to replace all spaces in a string with ‘%20’
/*
 Algorithm COMPRESS_STRING(str): 
 1.Get the input string str
 2.Find the length of the string
 3.Get the 0th element as lastChar
 4.Initialize the count as 1
 5.for i from 1 to len(str)
		currentChar=str[i]
		if(currentChar!=lastChar) then append lastChar and count to resultString
		else 
			count++
			lastChar=currentChar
 6.Finally append the lastChar and count to resultString - [edge case]
 7.return resultString
*/
package main

import (
	"fmt"
)

var sMap map[string]int
//Main function
func main() {
	str := " Helo Worl d! "
	fmt.Println("Input String :", str)
	fmt.Println(compressString(str))
}
//Function to compress the string
func compressString(str string) []string {
	length := len(str)
	if str == "" {
		return
	}
	var count int = 1
	lastChar := string([]rune(str)[0])
	for i := 1; i < length; i++ {
		ch := string([]rune(str)[i])
		if ch == " " {
			count++
		}
	}
	newLength := length + (2 * count)
	var sArr = make([]string, newLength)
	j := 0
	for i := 0; i < length; i++ {
		ch := string([]rune(str)[i])

		if ch == " " {
			sArr[j] = "%"
			j = j + 1
			sArr[j] = "2"
			j = j + 1
			sArr[j] = "0"
			j = j + 1
		} else {
			sArr[j] = ch
			j = j + 1
		}
	}
	return sArr
}
