package ascii

func HandleInput(data []string) (string, string) {
	str := ""
	font := "standard.txt"
	if len(data) == 1 {
		str = data[0]
		return str, font
	} else {
		for _, strin := range data {
			if strin == "shadow" || strin == "thinkertoy" || strin == "standard" {
				font = strin + ".txt"
				continue
			}
			str += strin + " "
		}
	}
	return str, font
}


