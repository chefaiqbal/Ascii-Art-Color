package ascii

func HandleInput(data []string) (string, string, string) {
	str := ""
	font := "standard.txt"
	color := "defult"
	if len(data) == 1 {
		str = data[0]
		return str, font, color
	} else {
		for _, strin := range data {
			if strin == "shadow" || strin == "thinkertoy" || strin == "standard" {
				font = strin + ".txt"
				continue
			}
			if strin[:2] == "--" && strin[:8] == "--color=" {
				color = strin[8:]
				continue
			}
			str += strin + " "
		}
	}
	return str, font, color
}
