package main

import (
	"fmt"
	"os"
	"strings"
	"parser/keywords"
	// "parser/expressions"
)

func main() {
	file := readFile("../../test.pr")
	lines := splitIntoLines(file)
	tokens := processLinesIntoTokens(lines)
	keywords := extractKeywords(tokens)
	fmt.Println(tokens)
	fmt.Println(keywords)
}

func extractKeywords(tokens [][]string) []string {
	var foundKeywords []string
	for _, line := range tokens {
		for _, token := range line {
			if keywords.Keywords.Contains(token) {
				foundKeywords = append(foundKeywords, token)
			}
		}
	}
	return foundKeywords
}

func processLinesIntoTokens(lines []string) [][]string {
	var allTokens [][]string
	for _, line := range lines {
		tokens := splitIntoTokens(line)
		allTokens = append(allTokens, tokens)
	}
	return allTokens
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
