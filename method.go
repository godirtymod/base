package base

import "strings"

// Apple returns a string about Apple for x times.
func Apple(x int) string {
	return strings.Repeat("This is A Apple 苹果 ", x)
}

// Banana is not banada.
func Banana() string {
	return "🍌 is new sexy"
}
