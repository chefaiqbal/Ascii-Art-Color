package ascii

import (
	"fmt"
	"strings"
)

func AsciiColorPrint(data string, font string, color string, toBeColored string) {
	fmt.Println(data, color)
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
			if word == toBeColored {
				for i := 0; i < 8; i++ {
					for _, letter := range word {
						res += GetLine(2+int(letter-' ')*9+i, font)
					}
					fmt.Println(GetColor(color), res, GetColor("reset"))
					res = ""
				}
			} else {
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
		}
	} else {
		sdata := strings.Split(data, " ")
		fmt.Println(sdata)
		for y := 0; y < len(sdata); y++ {
			if sdata[y] == toBeColored {
				for i := 0; i < 8; i++ {
					for _, let := range sdata {
						for _, letter := range let {
							res += GetLine(2+(int(letter)-32)*9+i, font)
						}
					}
					fmt.Println(GetColor(color), res, GetColor("reset"))
					res = ""
				}
			} else {
				for i := 0; i < 8; i++ {
					for _, letter := range data {
						res += GetLine(2+(int(letter)-32)*9+i, font)
					}
					fmt.Print(res)
					fmt.Println()
					res = ""
				}
			}
		}
	}
}
