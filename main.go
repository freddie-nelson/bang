package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// rainbow := [12]string{Red, Orange, Yellow, GreenYellow, Green, GreenCyan, Cyan, LightBlue, Blue, Purple, Pink, PinkRed}

	fmt.Println(Cyan + getLetter('A'))

	fmt.Println(Reset)
}

func getLetter(letter rune) string {
	data, err := ioutil.ReadFile("characters.txt")
	if err != nil {
		fmt.Println("Error while reading characters.txt", err)
		return ""
	}

	lines := strings.Split(string(data), "\n")
	lineNumber := int(letter-65) * 5
	var character []string = lines[lineNumber : lineNumber+5]

	return strings.Join(character, "\n")
}
