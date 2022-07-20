package wildcard

import (
	"strings"
)

const Wildcard = '*'

func Match(pattern, haystack string) bool {
	return matchStrings(pattern, haystack)
}

func MatchCI(pattern, haystack string) bool {
	return matchStrings(
		strings.ToLower(pattern),
		strings.ToLower(haystack),
	)
}

func MatchSlice(pattern string, haystack []string) bool {
	for _, value := range haystack {
		if Match(pattern, value) {
			return true
		}
	}

	return false
}

func MatchSliceCI(pattern string, haystack []string) bool {
	for _, value := range haystack {
		if MatchCI(pattern, value) {
			return true
		}
	}

	return false
}

func matchStrings(pattern, haystack string) bool {
	if len(pattern) == 0 {
		return len(haystack) == 0
	}

	if len(haystack) == 0 {
		return false
	}

	if string(Wildcard) == pattern {
		return true
	}

	if strings.ContainsRune(pattern, Wildcard) {

		return matchRunes([]rune(haystack), []rune(pattern), Wildcard)
	}

	return pattern == haystack

}

func matchRunes(haystack, pattern []rune, wildcard rune) bool {
	for len(pattern) > 0 {
		if wildcard == pattern[0] {
			return matchRunes(haystack, pattern[1:], wildcard) ||
				(len(haystack) > 0 && matchRunes(haystack[1:], pattern, wildcard))
		}

		if len(haystack) == 0 || haystack[0] != pattern[0] {
			return false
		}

		haystack = haystack[1:]
		pattern = pattern[1:]
	}

	return len(haystack) == 0 && len(pattern) == 0
}
