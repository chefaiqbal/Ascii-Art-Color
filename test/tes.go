package main

import "fmt"

func GetANSIColorCode(color string) string {
	codes := map[string]string{
		"black":   "\u001b[30m",
		"red":     "\u001b[31m",
		"green":   "\u001b[32m",
		"yellow":  "\u001b[33m",
		"blue":    "\u001b[34m",
		"magenta": "\u001b[35m",
		"cyan":    "\u001b[36m",
		"white":   "\u001b[37m",
		"reset":   "\u001b[0m",
	}
	return codes[color]
}

func main() {
	color := "blue"
	ansiCode := GetANSIColorCode(color)
	fmt.Println(ansiCode, "Hellow World", GetANSIColorCode("reset"))
}
