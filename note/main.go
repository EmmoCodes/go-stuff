package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {

	printSomething(1)
	printSomething(1.5)
	printSomething("hello")

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

	err = outputData(todo)

	if err != nil {
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}

	outputData(userNote)
}

func printSomething(value any) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("integer: ", intVal)
	}
	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("float: ", floatVal)
	}
	stringVal, ok := value.(string)
	if ok {
		fmt.Println("string: ", stringVal)
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println(" Saving the note succeeded.")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")

	content := getUserInput("Note content:")

	return title, content
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
