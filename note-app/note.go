package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const FILE_LOCATION = "notes.json"

var TOTAL_NOTES = 0

type Notes map[int]Note;

func getAllNotes()Notes {
	var notes Notes;
file,err:=os.Open(FILE_LOCATION);
 
   if ErrorHandling(err){
	panic(err)
   }
    defer file.Close()
   decoder:=json.NewDecoder(file);

   decoder.Decode(&notes);
   return notes;
   
}

func showAllNotes() {
	notes:=getAllNotes();
	
   for i,val := range notes{
	fmt.Println("-----------------------------------------")
		fmt.Printf("%d.Title:%s\nDescription:%s\n",i,val.Title,val.Description);
	fmt.Println("-----------------------------------------")
	}
     

}

func createANote() {
//note input

  note:=getNoteTitleAndDescription();

  notes :=getAllNotes();
  length :=len(notes);
  
  //append the note
  notes[length+1]=note;

//file operation
   file,err:= os.OpenFile(FILE_LOCATION, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644);
   
   if ErrorHandling(err){
	delete(notes,length+1);
	return
   }
   defer file.Close();
   encoder := json.NewEncoder(file);
   encoder.SetIndent(""," ");
   err2:= encoder.Encode(notes);

   if ErrorHandling(err2){
	delete(notes,length+1);
	return
   }

   fmt.Println("Successfully added note.")
  
   


}


//Error handling
func ErrorHandling(e error)bool{
  if e!=nil{
	fmt.Println(e);
	fmt.Println(errors.New("there was an error"))
	return true
  }
  return false
}
