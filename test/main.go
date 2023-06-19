package main

import (
	"fmt"
	"strings"
)

func colorizeText(data string, color string, toBeColored string) string {
	if strings.Contains(data, toBeColored) {
		coloredText := strings.ReplaceAll(data, toBeColored, fmt.Sprintf("\033[31m%s\033[0m", toBeColored))
		return coloredText
	}

	return data
}

func main() {
	data := "The quick brown fox jumps over the lazy dog"
	color := "red"
	toBeColored := "brown"

	result := colorizeText(data, color, toBeColored)
	fmt.Println(result)
}
