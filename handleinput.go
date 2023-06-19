package ascii

import (
	"fmt"
)

func findSmallestString(strings []string) string {
	smallest := strings[0] // Assume the first string is the smallest

	for _, str := range strings {
		if len(str) < len(smallest) {
			smallest = str
		}
	}

	return smallest
}

func HandleInput(data []string) (string, string, string, string) {
	str := ""
	font := "standard.txt"
	color := "defult"
	toBeColored := ""
	if len(data) == 1 {
		str = data[0]
		return str, font, color, toBeColored
	} else {
		toBeColored = findSmallestString(data)
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
	fmt.Printf("Print	:%s\nColor	:%s\nFont	:%s\nToColor	:%s", str, color, font, toBeColored)
	fmt.Println()
	return str, font, color, toBeColored
}
