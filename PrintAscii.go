package ascii

import (
	"fmt"
	"strings"
)

func GetIndex(data string, subdata string) int {
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
	// 				return count
	// 			}
	// 		}
	// 	}
	// }
	// return count
}

func AsciiPrint(data string, font string, color string, toBeColored string) {
	index := GetIndex(data, toBeColored)
	prev := 'a'
	severallines := false
	nline := 0
	for _, v := range data {
		if v == 'n' && prev == '\\' {
			nline++
			severallines = true
		}
		prev = v
	}
	// 3. Writing text line by line into res
	res := ""
	if severallines {
		args := strings.Split(data, "\\n")
		for _, word := range args {
			if word == "" {
				fmt.Println()
				continue
			}
			for i := 0; i < 8; i++ {
				for _, letter := range word {
					res += GetLine(2+int(letter-' ')*9+i, font)
				}
				fmt.Println(res)
				res = ""
			}
		}

	} else {
		for i := 0; i < 8; i++ {
			for x, letter := range data {
				if x >= index && x < index+len(toBeColored) {
					res += GetColor(color) + GetLine(2+(int(letter)-32)*9+i, font) + GetColor("reset")
				} else {
					res += GetLine(2+(int(letter)-32)*9+i, font)
				}
			}
			fmt.Println(res)
			res = ""

		}
	}
}
