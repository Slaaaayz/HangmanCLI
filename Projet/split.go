package hangman

func Splits(s, sep string) []string {
	for i := 0; i < len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			s = s[:i] + " " + s[i+len(sep):]
		}
	}
	return SplitWhiteSpaces(s)
}

func SplitWhiteSpaces(s string) []string {
	var tab []string
	b := false
	stock := ""
	for i := 0; i < len(s); i++ {
		if !(s[i] == 32 || s[i] == 9 || s[i] == '\n') {
			b = true
			stock += string(s[i])
			if i == len(s)-1 {
				tab = append(tab, stock)
			}
		} else if s[i] == 32 && b {
			b = false
			tab = append(tab, stock)
			stock = ""
		}
	}
	return tab
}
