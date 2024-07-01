package is

func Isnum(str string) bool {
	if str == ""{
		return false
	}
	
	out:= []rune(str)

	for _, c := range out{
		if c < '0' || c > '9'{
			return false
		}
	}

	return true
}
