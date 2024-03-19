package main

import (
	"fmt"
	"os"
	"strings"
	// "parser/keywords"
	// "parser/expressions"
)

func main() {
	file := readFile("../../test.pr")
	chars := splitIntoChars(file)
	fmt.Println(chars)
	
}





func splitIntoChars(file string) []string {
	lines := strings.Split(file, "")
	return lines
}

func readFile(filename string) string {
	result, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file.")
		os.Exit(1)
	}
	return string(result)
}
