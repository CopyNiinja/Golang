package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const FILE_LOCATION = "notes.json"

var TOTAL_NOTES = 0

func showAllNotes() {
	
   var notes map[int]Note;
  file,err:=os.Open(FILE_LOCATION);
 
   if ErrorHandling(err){
	return
   }
    defer file.Close()
   decoder:=json.NewDecoder(file);

   decoder.Decode(&notes);
   PrintNotes(notes)


}

func createANote() {

}
//Print Notes
func PrintNotes(notes map[int]Note){

	for i,val := range notes{
		fmt.Println(i,". ",val);
	}

}

//Error handling
func ErrorHandling(e error)bool{
  if e!=nil{
	fmt.Println(errors.New("There was an error!"))
	return true
  }
  return false
}
