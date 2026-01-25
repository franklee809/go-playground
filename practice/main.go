package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"exmaple.com/practice/note"
)

func main() {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content : ")
	// println(title, content)
	userNote, err := note.New(title, content)
	userNote.Display()
	println(err)

}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		println(err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
