package keywords

type StringList []string

var Keywords = StringList{"if", "for", "elif", "else", "while", "let", "const", "print", "puts"}

func (l StringList) Contains(word string) bool {
	for _, item := range l {
		if item == word {
			return true
		}
	}
	return false
}
