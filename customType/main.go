package main

//custom type (Opportunity to add methods to specific type?  Like prototype in js)

import (
	"errors"
	"fmt"
	"strconv"
)

//custom str type
type str string;

//custom methods to str type
   // print the string
func (s str) log(){
 fmt.Println(s);
}
  //length
func (s str) length()int{
	return len(s);
}		
  //ToNumber
func (s str ) toNumber()(int,error){
	value,err:=strconv.Atoi(string(s));
    if(err !=nil){
		return 0,errors.New("This string cant converted into int.")
	}
	
	return value,nil

}


//
type arr []int




func main (){


  //----------------CUSTOM TYPE STRING--------------------- 
	var x str;
     x="105";
	 x.log(); //"105"
	fmt.Println(x.length()) //3

	val,err := x.toNumber();
    
	fmt.Println(val)//105
	fmt.Printf("%T",val)//int
	fmt.Println(err)// <nill>
	//----------------------------------------------
       
	
}