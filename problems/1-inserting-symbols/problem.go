package insertingsymbols

func checkString(s string) bool {
	c := s[0]

	if s[len(s)-1] != c {
		return false
	}

	for i := 1; i < len(s)-1; i++ {
		if s[i] != c && s[i+1] != c {
			return false
		}
	}

	return true
}
