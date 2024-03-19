package main

import (
	"fmt"
	"strings"
	"os"
	"parser/keywords"
)

func main() {
	file := readFile("../../test.pr")
	lines := splitIntoLines(file)
	tokens := processLinesIntoTokens(lines)
	fmt.Println(tokens)
}

func processLinesIntoTokens(lines []string) [][]string {
	var keywords [][]string
	for _, line := range lines {
		tokens := splitIntoTokens(line)
		keywords = append(keywords, tokens)
	}
	return keywords
}

func findKeywords(tokens []string) []string {
	var result []string
	for _, token := range tokens {
		if keywords.Keywords.Contains(token) {
			result = append(result, token)
		}
	}
	return result
}

func splitIntoTokens(line string) []string {
	tokens := strings.Split(line, " ")
	return tokens
}

func splitIntoLines(file string) []string {
	lines := strings.Split(file, "\n")
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

