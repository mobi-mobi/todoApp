package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var todoList []string
var userInputMode int
var run bool = true
var fileurl string = "d:/SKOLA/GOVECI/todoApp/todoFile.txt"

func addMode() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter ToDo: ")
	input, _ := reader.ReadString('\n')

	addToDoElement(input)

}

func addToDoElement(text string) {
	todoList = append(todoList, text)

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
	for index, val := range todoList {
		fmt.Printf("\n%d %s", index+1, val)
	}
	fmt.Println("\n1-ADD   2-DELETE   3-MARK AS DONE   4-EXIT")
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

func openFile(url string) *os.File {
	file, err := os.OpenFile(url, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
	}

	return file
}

func readFile(f *os.File) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		todoText := scanner.Text()
		todoList = append(todoList, todoText+"\n")
	}
}

func rewriteFile() {
	todoString := strings.Join(todoList, "")
	todoBytes := []byte(todoString)
	err := os.WriteFile(fileurl, todoBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func markAsDone() {
	input := 0
	fmt.Print("Enter the order of the ToDo to mark it as done: ")
	fmt.Scanln(&input)
	todoList[input-1] = "TASK DONE!\n"
	/*
		todoList[input-1] = strings.Trim(todoList[input-1], "\n")
		todoList[input-1] += " DONE!\n"
	*/
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
		markAsDone()
		clearDisplay()
	case 4:
		run = false
	}

}

func main() {
	newFile := openFile(fileurl)
	readFile(newFile)
	for {
		printDisplay()
		handleModeInput()
		Modes()
		if !run {
			defer newFile.Close()
			break
		}
		rewriteFile()
	}
	clearDisplay()

}
