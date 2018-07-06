package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please provide one or more words to search")
}

// ParseLine recebe uma linha dos dados unicode do caracter e retorna seu
// código unicode e o nome do caracter.
func ParseLine(line string) (rune, string) {
	fields := strings.Split(line, ";")
	char, _ := strconv.ParseInt(fields[0], 16, 32)
	//fmt.Println(fields)

	return rune(char), fields[1]
}
