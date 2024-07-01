package is

func Isalpha(str string) bool {

	out := []rune(str)

	for c := range out{
		if !(c >= 'A' && c <= 'Z') && !(c >= 'a' && c <= 'z'){
			return false
		}
	}

	return true
	
}