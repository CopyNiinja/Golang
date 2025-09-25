package main

import "fmt"
var name = "faiyaz"
// name := "faiyaz"  (WRONG!!!!! only inside func) 
func main(){


	var num int; //default value : 0
	var str string; //default value : ""
	var isBoolean bool; //default value : false

	fmt.Printf("%d-%s-%t-%s",num,str,isBoolean,name); //  0 --false

	
}