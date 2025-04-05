package main

func findPairs(strings []string) int {
	oddGroups := make([]int, len(strings))
	evenGroups := make([]int, len(strings))

	for i := 0; i < len(strings); i++ {
		oddGroups[i] = -1
		evenGroups[i] = -1
	}

	count := 0

	for i := 0; i < len(strings); i++ {
		if oddGroups[i] == -1 {
			oddGroups[i] = i
		}

		if evenGroups[i] == -1 {
			evenGroups[i] = i
		}

		for j := i + 1; j < len(strings); j++ {
			var isOddEqual bool
			if oddGroups[i] == oddGroups[j] {
				isOddEqual = true
			} else if oddGroups[j] == -1 && oddGroups[i] != i {
				isOddEqual = false
			} else {
				isOddEqual = equalOdds(strings[i], strings[j])
			}

			var isEvenEqual bool
			if evenGroups[i] == evenGroups[j] {
				isEvenEqual = true
			} else if evenGroups[j] == -1 && evenGroups[i] != i {
				isEvenEqual = false
			} else {
				isEvenEqual = equalEvens(strings[i], strings[j])
			}

			if isOddEqual || isEvenEqual {
				count++
			}

			if isOddEqual && oddGroups[j] == -1 {
				oddGroups[j] = i
			}

			if isEvenEqual && evenGroups[j] == -1 {
				evenGroups[j] = i
			}
		}
	}

	return count
}

func equalOdds(s1, s2 string) bool {
	l := len(s1)

	if len(s1) != len(s2) {
		l := min(len(s1), len(s2))

		if l%2 == 0 {
			return false
		}

		if len(s1)-l > 1 || len(s2)-l > 1 {
			return false
		}
	}

	for i := 0; i < l; i += 2 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func equalEvens(s1, s2 string) bool {
	l := len(s1)

	if l < 2 {
		return false
	}

	if len(s1) != len(s2) {
		l = min(len(s1), len(s2))

		if l < 2 || l%2 != 0 {
			return false
		}

		if len(s1)-l > 1 || len(s2)-l > 1 {
			return false
		}
	}

	for i := 1; i < l; i += 2 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}
