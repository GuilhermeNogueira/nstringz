package string

// Reverse return a string in a reverse order tail recursivelly
func Reverse(s string) string {
	return reverse(s, "")
}

func reverse(s string, result string) string {
	b := []rune(s)
	if len(b) == 0 {
		return result
	}
	inverse := string(b[0]) + result
	return reverse(string(b[1:]), inverse)
}
