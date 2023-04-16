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
