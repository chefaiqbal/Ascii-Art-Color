package main

import (
	"fmt"
	"os"

	"ascii"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Println("Error: No enough arguments!!")
		return
	}
	data, font, color, toBeColored := ascii.HandleInput(args)
	if data == "" && font == "" && color == "" && toBeColored == "" {
		fmt.Println("Usage: go run . [OPTION] [STRING]\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		return
	} else {
		_, arerr := os.Stat("../fonts/" + font)
		if arerr != nil {
			fmt.Println("Error: ASCII-ART font file not found !")
			return
		}
		ascii.AsciiPrint(data, font, color, toBeColored)
	}
}
