package ft_strings

func Lowercase(str string) string {
	if str == ""{
		return ""
	}

	out := []rune(str)
	for i, c := range out{
		if c >= 'A' && c <= 'Z'{
			out[i] = c + 32 
		} 
	}
	return string(out)
}