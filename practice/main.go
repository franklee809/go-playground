package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"exmaple.com/practice/note"
	"exmaple.com/practice/todo"
)

func main() {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content : ")
	todoText := getUserInput("TODO text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	todo.Save()

	// println(title, content)
	userNote, err := note.New(title, content)
	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Println("Saving the note succeded!")

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
