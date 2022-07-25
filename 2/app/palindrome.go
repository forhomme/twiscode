package app

import "strings"

func Palindrome(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	start, maxLen, length := 0, 0, len(s)
	for i := 0; i < length; i++ {
		expandAroundCenter(s, i, i, &start, &maxLen)
		expandAroundCenter(s, i, i+1, &start, &maxLen)
	}
	return s[start : start+maxLen]
}

func expandAroundCenter(s string, left, right int, start, maxLen *int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left, right = left-1, right+1
	}
	if right-left-1 > *maxLen {
		*maxLen = right - left - 1
		*start = left + 1
	}
}
