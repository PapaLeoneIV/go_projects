package is

func Isascii(c rune) bool {
	return c >= 0 && c <= 127
}