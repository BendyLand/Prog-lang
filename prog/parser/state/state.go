package state

import (
	"fmt"
	"os"
	"regexp"
)

type LineParser struct {
	CurrentChar rune
	NextChar    rune
	Index       int
	State       string
}

func parseStringLiteral(line string) string {
	pattern, err := regexp.Compile("\".*\"")
	if err != nil {
		fmt.Println("There was a problems creating the regex")
		os.Exit(1)
	}
	return pattern.FindString(line)
}

func (lp *LineParser) Parse(line string) []string {
	var tokens []string
	token := ""
	Loop: for i := range len(line) - 1 {
		lp.UpdateState(i, rune(line[i]), rune(line[i+1]))
		fmt.Println(lp.State, token)
		switch lp.State {
		case "Parser":
			if line[i] == ' ' {
				tokens = append(tokens, token)
				token = ""
			} else {
				token += string(line[i])
			}
		case "StringLiteral":
			token += string(line[i])
		case "StringLiteralTerminator":
			tokens = append(tokens, token)
		case "Comment":
			continue Loop
		}
	}
	return tokens
}

func (lp *LineParser) UpdateState(index int, current, next rune) {
	lp.Index = index
	lp.CurrentChar = current
	lp.NextChar = next
	switch {
	case lp.CurrentChar == '"' && lp.State != "StringLiteral":
		lp.State = "StringLiteral"
	case lp.CurrentChar == '"' && lp.State == "StringLiteral":
		lp.State = "StringLiteralTerminator"
	case lp.State == "StringLiteralTerminator":
		lp.State = "Parser"
	case lp.CurrentChar == '/' && lp.NextChar == '/':
		lp.State = "Comment"
	case lp.State == "Comment" && lp.NextChar == '\n':
		lp.State = "Parser"
	}

}
