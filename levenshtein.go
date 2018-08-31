package woma

// Levenshtein distance between two given strings.
// see https://gist.github.com/andrei-m/982927#gistcomment-1931258
// and https://jsperf.com/levenshtein123456
func Levenshtein(word1, word2 string) int {
	return levenshtein([]rune(word1), []rune(word2))
}

func levenshtein(a, b []rune) int {
	if len(a) == 0 {
		return len(b)
	}
	if len(b) == 0 {
		return len(a)
	}

	// swap to save some memory O(min(a,b)) instead of O(a)
	if len(a) > len(b) {
		a, b = b, a
	}

	// init the row
	row := make([]int, len(a)+1)
	for j := 1; j <= len(a); j++ {
		row[j] = j
	}

	for i := 1; i <= len(b); i++ {
		prev := i
		var val int
		for j := 1; j <= len(a); j++ {
			if b[i-1] == a[j-1] {
				val = row[j-1]
			} else {
				val = min(row[j-1]+1, // substitution
					min(prev+1, // insertion
						row[j]+1)) // deletion
			}
			row[j-1] = prev
			prev = val
		}
		row[len(a)] = prev
	}
	return row[len(a)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
