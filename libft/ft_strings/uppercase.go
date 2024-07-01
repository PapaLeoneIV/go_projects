package ft_strings

func Uppercase(str string) string {
	if str == ""{
		return ""
	}

	out := []rune(str)
	for i, c := range out{
		if c >= 'a' && c <= 'z'{
			out[i] = c - 32 
		} 
	}
	return string(out)
}