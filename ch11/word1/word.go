// Package word provides utilities for word games.
package word

import "unicode"

// IsPalindrome reports wether s reads the same forward and backward.
// Letter case is ignored, as are non-letters.
// (Our first attempt.)
func IsPalindrome(s string) bool {
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
