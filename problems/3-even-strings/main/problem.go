package main

func findPairs(strings []string) int {
	count := 0

	for i := 0; i < len(strings); i++ {
		for j := i + 1; j < len(strings); j++ {
			if equalOdds(strings[i], strings[j]) || equalEvens(strings[i], strings[j]) {
				count++
			}
		}
	}

	return count
}

func equalOdds(s1, s2 string) bool {
	l := min(len(s1), len(s2))

	if l%2 == 0 && len(s1) != len(s2) {
		return false
	}

	if len(s1)-l > 1 || len(s2)-l > 1 {
		return false
	}

	for i := 0; i < l; i += 2 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func equalEvens(s1, s2 string) bool {
	l := min(len(s1), len(s2))

	if l < 2 {
		return false
	}

	if l%2 != 0 && len(s1) != len(s2) {
		return false
	}

	if len(s1)-l > 1 || len(s2)-l > 1 {
		return false
	}

	for i := 1; i < len(s1); i += 2 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
