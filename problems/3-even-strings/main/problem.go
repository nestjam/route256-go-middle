package main

import (
	"hash/maphash"
)

func findPairs(strings []string) int {
	oddGroups := make([]int, len(strings))
	evenGroups := make([]int, len(strings))

	h := &maphash.Hash{}
	oddHashes := make([]uint64, len(strings))
	evenHashes := make([]uint64, len(strings))

	for i := 0; i < len(strings); i++ {
		oddGroups[i] = -1
		evenGroups[i] = -1
		oddHashes[i] = hashOdds(h, strings[i])
		evenHashes[i] = hashEvens(h, strings[i])
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
			} else if oddHashes[i] == oddHashes[j] {
				isOddEqual = equalOdds(strings[i], strings[j])
			}

			var isEvenEqual bool
			if evenGroups[i] == evenGroups[j] {
				isEvenEqual = true
			} else if evenGroups[j] == -1 && evenGroups[i] != i {
				isEvenEqual = false
			} else if evenHashes[i] == evenHashes[j] {
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
	n := min(len(s1), len(s2))

	if n%2 == 0 {
		n--
	}

	if len(s1)-n > 1 || len(s2)-n > 1 {
		return false
	}

	l, r := 0, n-1
	for i := 0; i < max(n/2, 1); i++ {
		if s1[l] != s2[l] {
			return false
		}

		if s1[r] != s2[r] {
			return false
		}

		l += 2
		r -= 2
	}

	return true
}

func equalEvens(s1, s2 string) bool {
	n := min(len(s1), len(s2))

	if n%2 == 1 {
		n--
	}

	if n < 2 {
		return false
	}

	if len(s1)-n > 1 || len(s2)-n > 1 {
		return false
	}

	l, r := 1, n-1
	for i := 0; i < max(n/2, 1); i++ {
		if s1[l] != s2[l] {
			return false
		}

		if s1[r] != s2[r] {
			return false
		}

		l += 2
		r -= 2
	}

	return true
}

func hashOdds(h *maphash.Hash, s string) uint64 {
	h.Reset()

	for i := 0; i < len(s); i += 2 {
		h.WriteByte(s[i])
	}

	return h.Sum64()
}


func hashEvens(h *maphash.Hash, s string) uint64 {
	h.Reset()

	for i := 1; i < len(s); i += 2 {
		h.WriteByte(s[i])
	}

	return h.Sum64()
}