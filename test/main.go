package main

import (
	"fmt"
	"strings"
)

func FindSubstring(strings []string) string {
	for _, substr := range strings {
		found := false
		for _, str := range strings {
			if str != substr && containsString(str, substr) {
				found = true
				break
			}
		}
		if !found {
			return substr
		}
	}
	return "-"
}

func containsString(str, substr string) bool {
	return strings.Contains(str, substr)
}

func main() {
	input2 := []string{"hello", "ahmed Mohamed", "isa hollow", "ahmed", "isa Mohamed", "yellow world"}
	substring2 := toarray(input2)
	for i := 0; i < len(substring2); i++ {
		fmt.Print(substring2[i], "--") // Output: isa
	}
	input3 := []string{"hellow world", "isa"}
	substring3 := toarray(input3)
	fmt.Println(substring3) // Output: -
}

func toarray(data []string) []string {
	var res []string
	for _, word := range data {
		if strings.Contains(word, " ") {
			temp := strings.Split(word, " ")
			res = append(res, temp...)
		} else {
			res = append(res, word)
		}
	}
	return res
}
