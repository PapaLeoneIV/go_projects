package ft_strings

func Strlen(str string) int {

	if str == ""{
		return 0
	}

	length := 0

	for range str {
		length++
	}

	return length
}