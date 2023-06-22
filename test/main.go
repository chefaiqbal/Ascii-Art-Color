package main

import (
	"fmt"
	"strings"
)

func ind(data string, subdata string) int {
	if strings.Contains(data, subdata) {
		index := strings.Index(data, subdata)
		return index
	} else {
		return -1
	}
	// count := 0
	// for i := 0; i < len(data); i++ {
	// 	if strings.Contains(data, subdata) {
	// 		for _, char := range data {
	// 			count++
	// 			if char == rune(subdata[0]) {
	// 				return count - 1
	// 			}
	// 		}
	// 	}
	// }
	// return count
}

func arr(s1 string, s2 string) {
	index := ind(s1, s2)
	fmt.Println(index)
	fmt.Println(s1)
	for i := 0; i < len(s1); i++ {
		if i >= index && i < index+len(s2) {
			fmt.Print("^")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {
	str := "Hellow World!"
	substr := "ld"
	arr(str, substr)
}
