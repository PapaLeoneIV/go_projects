package to

import (
	"libft/ft_strings"
)

func Atoi(str string) (int,  string) {
	res
	isneg := false
	offset := 0

	if str == "" {
		return 0, "errore"
	}

	out := []byte(str)

	if out[0] == '+' || out[0] == '-' {
		offset = 1
		if out[0] == '-' {
			isneg = true
		}
	}

	for start := offset; start < ft_strings.Strlen(str); start++ {
		if out[start] < '0' || out[start] > '9' {
			return 0, "invalid input"
		}
		res = res*10 + int(out[start]-48)
	}

	if isneg {
		res = -res
	}

	return res, ""

}
