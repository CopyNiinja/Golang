package main

import (
	"fmt"
	"time"
)

//imports


func main() {
 

	for{
		var choice int;
        //
		PrintOptions();
       choice=getUserChoice("Your choice:")

	   //
	   switch choice{
	   case 1:
		showAllNotes();
	   case 2:
		createANote();
	   default:
		fmt.Println("Thank you for using our app.")
		return 
	   }

 	 
	}
}

type Note struct{
	Title string
	Description string
	CreatedAt time.Time
}

func newNote(title ,description string)Note{
	return Note{
		Title: title,
		Description: description,
		CreatedAt: time.Now(),
	}
}

func PrintOptions (){
	
	fmt.Println("Welcome to Note App")
	fmt.Println("Choose an option to continue:")
	fmt.Println("1.Show notes")
	fmt.Println("2.Create new note")
	fmt.Println("3.Exist")
}