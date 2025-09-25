package main

import "fmt"

func main(){
     
    //var Array_Name = [length]type {array_element1,array_element2,...}
 	 
	//declaration 

	var arr1 = [10]int{1,2,3};
	arr2 := [10]int{};
	//size inferred from values
	arr3:=[...]string {"Hello","world","hehehe"};

	//array specific index element declare

	arr4 := [10] int {4:707,8:999}; //[0,0,0,0,707,0,0,0,999,0]
	
	//array item modification
    // Array_Name[index]=value
	arr4[2]=33



// ------------------print----------------------------
	fmt.Println(len (arr1)) // 10
	fmt.Println(arr1) // [1,2,3,0,0,0,0,0,0,0]
	fmt.Println(arr2)// [1,2,3,0,0,0,0,0,0,0]
	fmt.Println(arr3)// [Hello world hehehe]
	fmt.Println(arr4)//[0,0,0,0,707,0,0,0,999,0]



}