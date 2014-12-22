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
