package ascii

import (
	"bufio"
	"os"
	"strings"
)

// getting the line data
func GetLine(num int, font string) string {
	f, _ := os.Open("../fonts/" + font)
	defer f.Close()

	content := bufio.NewReader(f)
	for i := 0; i < num; i++ {
		line, _ := content.ReadString('\n')

		// Trim any leading or trailing whitespace
		line = strings.TrimSuffix(line, "\n")
		if i == num-1 {
			return line
		}
	}

	return ""
}
