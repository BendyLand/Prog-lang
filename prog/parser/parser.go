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
	chars := splitIntoChars(file)
	fmt.Println(chars)
	
}





































func matchKeywordsToTokens(keywords []string, tokens [][]string) map[string][]string {
	resultMap := make(map[string][]string, len(keywords))
	var filteredTokens [][]string
	for _, keyword := range keywords {
		temp := filterContainsKeyword(tokens, keyword)
		filteredTokens = append(filteredTokens, temp)
	}
	for i := 0; i < len(keywords); i++ {
		resultMap[keywords[i]] = filteredTokens[i]
	}
	return resultMap
}

func filterContainsKeyword(tokens [][]string, keyword string) []string {
	var newSlice []string
	for _, line := range tokens {
		for _, word := range line {
			if word == keyword {
				newLine := strings.Join(line, " ")
				newSlice = append(newSlice, newLine)
			}
		}
	}
	return newSlice
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
