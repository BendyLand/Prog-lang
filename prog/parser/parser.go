package main

import (
	"fmt"
	"os"
	"parser/keywords"
	"parser/state"
	"strings"
)

func main() {
	file := readFile("../../test.pr")
	linesInTokens := parse(file)
	for _, line := range linesInTokens {
		for _, token := range line {
			switch token {
			case "print":
				parsedLine, err := keywords.ParsePrintFn(line)
				if err == nil {
					fmt.Print(parsedLine) // no newline!
				} else {
					fmt.Println("Error here:", err)
				}
			case "puts":
				parsedLine, err := keywords.ParsePutsFn(line)
				if err == nil {
					fmt.Println(parsedLine) // yes newline!
				} else {
					fmt.Println("Error here:", err)
				}
			}
		}
	}
}

// this is the main parser, which will possess context aware characteristics
func parse(file string) [][]string {
	var result [][]string
	lines := strings.Split(file, "\n")
	parser := state.LineParser{
		Index: 0,
		State: "Parser",
	}
	for _, line := range lines {
		tokens := parser.Parse(line) // context state is updated here
		result = append(result, tokens)
	}
	return result
}

func readFile(filename string) string {
	result, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file.")
		os.Exit(1)
	}
	return string(result)
}
