func lengthOfLastWord(s string) int {

	var length int
	wordTrimmed := strings.Split(strings.Trim(s, " "), "")

	for i := len(wordTrimmed) - 1; i >= 0; i-- {

		if wordTrimmed[i] != " " {
			length = length + 1
		} else {
			return length
		}
	}
	return length
}
