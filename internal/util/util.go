package util

import (
	"strings"

	"github.com/mattn/go-runewidth"
)

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Width(s string) int {
	return runewidth.StringWidth(s)
}

func Pad(s string, length int) string {
	w := Width(s)
	if w < length {
		return s + strings.Repeat(" ", length-w)
	}
	return s
}

func String(s string) *string {
	return &s
}

func Bool(b bool) *bool {
	return &b
}

func Strings(ss []string) *[]string {
	return &ss
}
