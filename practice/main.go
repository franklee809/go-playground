package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"exmaple.com/practice/note"
	"exmaple.com/practice/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

// type outputtable interface {
// 	Save() error
// 	Display()
// }

func main() {
	printSomething(1)
	printSomething("sample text")
	printSomething(true)
	title := getUserInput("Note title: ")

	content := getUserInput("Note content : ")
	todoText := getUserInput("TODO text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		return
	}

	outputData(todo)

	// println(title, content)
	userNote, err := note.New(title, content)
	err = outputData(userNote)

	if err != nil {
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

func printSomething(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Boolean:", v)
	default:
		fmt.Println("Unknown type")
	}
	fmt.Println("Hello, World!")
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed. ")
		return err

	}
	fmt.Println("Saving the note succeeded! ")
	return nil
}
