package ascii

import (
	"fmt"
	"strings"
)

func HandleInput(data []string) (string, string, string, string) {
	str := ""
	font := "standard.txt"
	color := "defult"
	toBeColored := ""
	if len(data) == 1 {
		str = data[0]
		return str, font, color, toBeColored
	} else {
		toBeColored = FindSubstring(data)
		for _, strin := range data {
			if strin == "shadow" || strin == "thinkertoy" || strin == "standard" {
				font = strin + ".txt"
				continue
			}
			if strin[:2] == "--" && strin[:8] == "--color=" {
				color = strin[8:]
				continue
			}
			if strin == toBeColored {
				continue
			}
			str += strin + " "
		}
	}
	fmt.Printf("Data	:%v\nFont	:%v\nColor	:%v\nToBeColo:%v\n", str, font, color, toBeColored)
	return str, font, color, toBeColored
}

func FindSubstring(data []string) string {
	prev := ""
	found := false
	fmt.Println(data, len(data))

	for _, str := range data {
		if str[:2] != "--" {
			if strings.Contains(str, prev) && !found {
				found = true
				fmt.Println(prev)
				return prev
			} else if found {
				prev = str
				found = false
			}
		}
	}
	if found {
		fmt.Println(prev)
		return prev
	} else {
		fmt.Println(prev)
		return ""
	}
}
