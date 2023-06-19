package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func PrintASCIIArt(character string) {
	// Read the contents of the thinkertoy.txt file
	content, err := ioutil.ReadFile("../fonts/thinkertoy.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Split the content into lines
	lines := strings.Split(string(content), "\n")

	// Create a map to store the character-to-ASCII-art mapping
	charMap := make(map[string][]string)

	// Iterate over the lines and populate the charMap dictionary
	for i := 0; i < len(lines); i += 9 {
		char := lines[i]
		art := lines[i+1 : i+9]
		charMap[char] = art
	}

	// Check if the character exists in the charMap
	art, exists := charMap[character]
	if exists {
		// Print the ASCII art for the character
		fmt.Println(strings.Join(art, "\n"))
	} else {
		fmt.Println("Character not found in the ASCII art database.")
	}
}

func main() {
	PrintASCIIArt("A") // Prints ASCII art for the character "A"
	PrintASCIIArt("B") // Prints ASCII art for the character "B"
	PrintASCIIArt("$") // Prints "Character not found in the ASCII art database."
}
