package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	message := "happy new year"
	messageLength := len(message)

	var lines [5][]string

	lines[0] = make([]string, 0, messageLength)
	lines[1] = make([]string, 0, messageLength)
	lines[2] = make([]string, 0, messageLength)
	lines[3] = make([]string, 0, messageLength)
	lines[4] = make([]string, 0, messageLength)

	createRows(&lines, message)
	printRows(lines, messageLength)
	fmt.Println(Reset)
}

func createRows(rows *[5][]string, message string) {
	for _, v := range strings.ToUpper(message) {
		letter := getLetter(v)
		if letter == nil {
			return
		}

		rows[0] = append(rows[0], letter[0])
		rows[1] = append(rows[1], letter[1])
		rows[2] = append(rows[2], letter[2])
		rows[3] = append(rows[3], letter[3])
		rows[4] = append(rows[4], letter[4])
	}
}

func printRows(rows [5][]string, messageLength int) {
	rainbow := [12]string{Red, Orange, Yellow, GreenYellow, Green, GreenCyan, Cyan, LightBlue, Blue, Purple, Pink, PinkRed}
	var lines [5]string

	column := 0

	for i := 0; i < messageLength; i++ {
		for j := 0; j < len(lines); j++ {
			var s string

			for _, char := range rows[j][i] {
				s = s + rainbow[column/2%len(rainbow)] + string(char)
				column++
			}

			lines[j] = fmt.Sprintf("%s %s", lines[j], strings.TrimRight(s, "\r"))

			if j != len(lines)-1 {
				column = 0 + i*len(rows[j][i])
			}
		}
	}

	fmt.Println(strings.Join(lines[:], "\n"))
}

func getLetter(letter rune) []string {
	data, err := ioutil.ReadFile("characters.txt")
	if err != nil {
		fmt.Println("Error while reading characters.txt", err)
		return nil
	}

	lines := strings.Split(string(data), "\n")
	lineNumber := int(letter-65) * 5

	if lineNumber < 0 {
		return []string{
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
		}
	}

	var character []string = lines[lineNumber : lineNumber+5]

	return character
}
