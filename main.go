package main

import (
	"fmt"
	"strings"
)

func IsAwesome(text string) bool {
	return strings.EqualFold("awesome", text)
}

func main() {
	fmt.Printf("Is 'awesome', AWESOME? %v\n", IsAwesome("awesome"))
}
