package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-notes/note"
	"example.com/go-notes/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func printSomething(value any) {

	typedVal, ok := value.(int)

	if ok {
		fmt.Println("Integer:", typedVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println("Float:", floatVal)
		return
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println(stringVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println(value)
	// }
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()
	err = data.Save()
	if err != nil {
		fmt.Println("Saving the note failed!")
		return err
	}

	fmt.Println("Saving the note succeeded!")
	return nil
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
