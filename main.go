package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type todoElement struct {
	decor string
	Text  string
}

var todoList []todoElement
var userInputMode int
var run bool = true

func addMode() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter ToDo: ")
	input, _ := reader.ReadString('\n')

	addToDoElement(input)

}

func addToDoElement(text string) {
	todoList = append(todoList, todoElement{decor: "* ", Text: text})
}

func removeMode() {
	input := 0
	fmt.Print("Enter the order of the ToDo to delete it: ")
	fmt.Scanln(&input)
	todoList[input-1] = todoList[len(todoList)-1]
	todoList = todoList[:len(todoList)-1]
}

func printDisplay() {
	fmt.Println("TODO LIST")
	fmt.Println("----------")
	for _, val := range todoList {
		fmt.Println(val.decor, val.Text)
	}
	fmt.Println("\n1-ADD   2-DELETE   3-EXIT")
}

func handleModeInput() {
	fmt.Print("Type the number of the operation you wish to perform: ")
	fmt.Scanln(&userInputMode)
}

func clearDisplay() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Modes() {

	switch userInputMode {
	case 1:
		addMode()
		clearDisplay()
	case 2:
		removeMode()
		clearDisplay()
	case 3:
		run = false
	}

}

func main() {
	for {
		printDisplay()
		handleModeInput()
		Modes()
		if !run {
			break
		}
	}
	clearDisplay()

}
