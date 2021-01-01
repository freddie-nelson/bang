package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(fmt.Errorf("%sError: please enter a string to bannerified%s", Red, Reset))
		return
	} else if len(os.Args) > 3 {
		fmt.Println(fmt.Errorf("%sError: too many arguments%s", Red, Reset))
		return
	}

	var (
		message        string
		charHeight     int
		isCharacters3D bool
	)

	if strings.ToUpper(os.Args[1]) == "--2D" {
		message = os.Args[2]
		charHeight = 5
		isCharacters3D = false

	} else {
		message = os.Args[1]
		charHeight = 6
		isCharacters3D = true
	}

	DisplayBanner(message, charHeight, isCharacters3D)
}

// DisplayBanner : displays a large ASCII text banner of the given message in a rainbow color
func DisplayBanner(message string, charHeight int, isCharacters3D bool) {
	messageLength := len(message)

	lines := make([][]string, charHeight)
	for i := 0; i < charHeight; i++ {
		lines[i] = make([]string, 0, messageLength)
	}

	lines = createRows(lines, message, charHeight, isCharacters3D)
	printRows(lines, messageLength, charHeight)
	fmt.Println(Reset)
}

func createRows(rows [][]string, message string, charHeight int, isCharacters3D bool) [][]string {
	for _, char := range strings.ToUpper(message) {
		letter := getLetter(char, charHeight, isCharacters3D)
		if letter == nil {
			return nil
		}

		for i := 0; i < charHeight; i++ {
			rows[i] = append(rows[i], letter[i])
		}
	}

	return rows
}

func printRows(rows [][]string, messageLength int, charHeight int) {
	rainbow := [12]string{Red, Orange, Yellow, GreenYellow, Green, GreenCyan, Cyan, LightBlue, Blue, Purple, Pink, PinkRed}
	lines := make([]string, charHeight)
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

			lines[j] = fmt.Sprintf("%s %s", lines[j], strings.TrimRight(s, "\r\n"))
		}

		lastCol = column
	}

	fmt.Println("\n" + strings.Join(lines[:], "\n"))
}

func getLetter(letter rune, charHeight int, isCharacters3D bool) []string {
	var lines []string
	if isCharacters3D {
		lines = strings.Split(Characters3D, "\n")
	} else {
		lines = strings.Split(Characters2D, "\n")
	}

	lineNumber := int(letter-65) * charHeight

	if lineNumber < 0 {
		space := make([]string, charHeight)
		for i := range space {
			space[i] = "      "
		}

		return space[:]
	}

	var character []string = lines[lineNumber : lineNumber+charHeight]

	return character
}
