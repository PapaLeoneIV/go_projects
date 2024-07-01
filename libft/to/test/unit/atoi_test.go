package to

import (
	"testing"
	"libft/to/atoi"
)

func TestAtoi(t *testing.T){

	testcase:= []struct{
			origin string
			output int
			expectedmsg string

		}{
		{"123", 123, ""},
		{"-123", -123, ""},
        {"", 0, "empty string"},
        {"abc", 0, "invalid input"},
        {"0123", 123, ""},
        {"-0", 0, ""},
        {"+123", 123, ""},
        {"123abc", 0, "invalid input"}, // Depends on function's strictness
    }
	
	for _, elem := range testcase {
		res  , msg := to.Atoi(elem.origin)
		if (res != elem.output || msg != elem.expectedmsg){

			t.Errorf("Atoi(%q) = %d, %q; expected %d, %q", elem.origin, res, msg, elem.output, elem.expectedmsg)	
		}
	}
}
