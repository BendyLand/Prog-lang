package keywords

type StringList []string

var Keywords = StringList{"if", "for", "elif", "else", "while", "let", "const", "print", "puts", "let", "mut"}

func CheckAgainstKeywords(word string) bool {
	return Keywords.contains(word)
}

func IsComment(token string) bool {
	return token == "//"
}

func (l StringList) contains(word string) bool {
	for _, item := range l {
		if item == word {
			return true
		}
	}
	return false
}
