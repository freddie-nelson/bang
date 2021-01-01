package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const charHeight int = 6
const charactersFile string = "3dcharacters.txt"

func main() {

	if len(os.Args) < 2 {
		fmt.Println(fmt.Errorf("%sError: please enter a string to bannerified%s", Red, Reset))
		return
	} else if len(os.Args) > 2 {
		fmt.Println(fmt.Errorf("%sError: too many arguments%s", Red, Reset))
		return
	}

	message := os.Args[1]
	messageLength := len(message)

	var lines [charHeight][]string

	for i := 0; i < charHeight; i++ {
		lines[i] = make([]string, 0, messageLength)
	}

	createRows(&lines, message)
	printRows(lines, messageLength)
	fmt.Println(Reset)
}

func createRows(rows *[charHeight][]string, message string) {
	for _, v := range strings.ToUpper(message) {
		letter := getLetter(v)
		if letter == nil {
			return
		}

		for i := 0; i < charHeight; i++ {
			rows[i] = append(rows[i], letter[i])
		}
	}
}

func printRows(rows [charHeight][]string, messageLength int) {
	rainbow := [12]string{Red, Orange, Yellow, GreenYellow, Green, GreenCyan, Cyan, LightBlue, Blue, Purple, Pink, PinkRed}
	var lines [charHeight]string

	column := 0
	lastCol := 0

	for i := 0; i < messageLength; i++ {
		for j := 0; j < charHeight; j++ {
			column = lastCol
			var s string

			for _, char := range rows[j][i] {
				s = s + rainbow[column%len(rainbow)] + string(char)
				column++
			}

			lines[j] = fmt.Sprintf("%s %s", lines[j], strings.TrimRight(s, "\r"))
		}

		lastCol = column
	}

	fmt.Println("\n" + strings.Join(lines[:], "\n"))
}

func getLetter(letter rune) []string {
	data, err := ioutil.ReadFile(charactersFile)
	if err != nil {
		fmt.Println("Error while reading characters file", err)
		return nil
	}

	lines := strings.Split(string(data), "\n")
	lineNumber := int(letter-65) * charHeight

	if lineNumber < 0 {
		var space [charHeight]string
		for i := range space {
			space[i] = "      "
		}

		return space[:]
	}

	var character []string = lines[lineNumber : lineNumber+charHeight]

	return character
}
