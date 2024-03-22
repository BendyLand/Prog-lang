package keywords

import (
	"fmt"
)

func ParsePutsFn(tokens []string) (string, error) {
	for i, token := range tokens {
		if i > 1 {
			continue
		} // temporary handling of extra tokens; will need to refactor to handle string interpolation
		switch {
		case i == 0 && token != "puts":
			return "", fmt.Errorf("Error: Not a `puts` function.\n")
		case i != 0 && token[0] != '"' && token[len(tokens)-1] != '"':
			return "", fmt.Errorf("Error: Argument not a string\n")
		}
	}
	return tokens[1], nil
}

func ParsePrintFn(tokens []string) (string, error) {
	for i, token := range tokens {
		if i > 1 {
			continue
		} // same as line 11; temporary fix
		switch {
		case i == 0 && token != "print":
			return "", fmt.Errorf("Error: Not a `print` function.\n")
		case i != 0 && token[0] != '"' && token[len(tokens)-1] != '"':
			return "", fmt.Errorf("Error: Argument not a string.\n")
		}
	}
	return tokens[1], nil
}
