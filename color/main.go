package main

import (
	"fmt"
	"os"

	ascii "simple/main.go"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 0 {
		fmt.Println("Please provide enough arguments !!")
		return
	}
	data, font, color, toBeColored := ascii.HandleInput(args)
	_, arerr := os.Stat("../fonts/" + font)
	if arerr != nil {
		fmt.Println("Error: ASCII-ART font file not found !")
		return
	}
	ascii.AsciiPrint(data, font, color, toBeColored)
}
