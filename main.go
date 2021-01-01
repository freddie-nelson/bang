package main

import (
	"fmt"
	"io/ioutil"
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
		charactersFile string
	)

	if os.Args[1] == "--3d" {
		message = os.Args[2]
		charHeight = 6
		charactersFile = "3dcharacters.txt"

	} else {
		message = os.Args[1]
		charHeight = 5
		charactersFile = "characters.txt"
	}

	DisplayBanner(message, charHeight, charactersFile)
}

// DisplayBanner : displays a large ASCII text banner of the given message in a rainbow color
func DisplayBanner(message string, charHeight int, charactersFile string) {
	messageLength := len(message)

	lines := make([][]string, charHeight)
	for i := 0; i < charHeight; i++ {
		lines[i] = make([]string, 0, messageLength)
	}

	lines = createRows(lines, message, charHeight, charactersFile)
	printRows(lines, messageLength, charHeight)
	fmt.Println(Reset)
}

func createRows(rows [][]string, message string, charHeight int, charactersFile string) [][]string {
	for _, char := range strings.ToUpper(message) {
		letter := getLetter(char, charHeight, charactersFile)
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

func getLetter(letter rune, charHeight int, charactersFile string) []string {
	data, err := ioutil.ReadFile(charactersFile)
	if err != nil {
		fmt.Println("Error while reading characters file", err)
		return nil
	}

	lines := strings.Split(string(data), "\n")
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
