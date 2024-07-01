package main

import (
	"fmt"
	"libft/to" 
)

func main() {
	testString := make([]string, 2)

	testString[0] = "1213"

	numRes, err := to.Atoi(testString[0])

	if err != "" {
		fmt.Println("Error:", numRes)
		return
	}

	fmt.Println("Numeric:", numRes)
}
