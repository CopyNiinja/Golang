package main

import (
	"bufio"
	"fmt"
	"os"
)

//get note  title and description

func getNoteTitleAndDescription() Note {
	title := getUserInput("The title:")
	description := getUserInput("The description:")

	return newNote(title, description)
}

//get user input

func getUserInput(prompt string) string {
	//variables
	reader := bufio.NewReader(os.Stdin)
	var input string
	fmt.Println(prompt)

	input, _ = reader.ReadString('\n')

	return input

}
