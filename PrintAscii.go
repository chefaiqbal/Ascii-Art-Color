package ascii

import (
	"fmt"
	"strings"
)

func AsciiPrint(data string, font string, color string, toBeColored string) {
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
			for _, letter := range data {
				res += GetLine(2+(int(letter)-32)*9+i, font)
			}
			if color != "" {
				fmt.Println(GetColor(color), res, GetColor("defult"))
			} else {
				fmt.Println(res)
			}
			res = ""
		}
	}
}
