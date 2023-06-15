package ascii

import "strings"

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
	return str, font, color, toBeColored
}

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
