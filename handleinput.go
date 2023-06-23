package ascii

import (
	"fmt"
	"strings"
)

func FindSubString(data []string) string {
	smallest := data[0] // Assume the first string is the smallest

	for _, str := range data {
		if len(str) < len(smallest) {
			smallest = str
		}
	}
	for _, str := range data {
		if str != smallest && strings.Contains(str, smallest) {
			return smallest
		}
	}
	return ""
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
		toBeColored = FindSubString(data)
		for _, strin := range data {
			if strin == toBeColored {
				continue
			}
			if strin[:2] == "--" {
				if len(strin) > 8 && strin[:8] == "--color=" {
					color = strin[8:]
					ok := IsColorSupported(color)
					if !ok {
						fmt.Println("Error: Coloris not supported!\nSuported colors are:\n-black	-red\n-green	-yellow\n-blue	-magenta\n-cyan	-white\n-orange	-reset")
						return "", "", "", ""
					} else {
						continue
					}
				} else {
					return "", "", "", ""
				}
			}
			if strin == "shadow" || strin == "thinkertoy" || strin == "standard" {
				font = strin + ".txt"
				continue
			}

			str += strin + " "
		}
	}
	return str, font, color, toBeColored
}
