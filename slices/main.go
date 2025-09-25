package main

import "fmt"

// len() function - returns the length of the slice (the number of elements in the slice)
// cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink)

func main(){
    // declaration
	var slice1 = []int {};
    slice2 :=[]string { "Hello","world"};
    //make function
	var slice4 = make([]int,10,20); // make([]type,length,capacity)
	
	fmt.Println(slice4) //[0,0,0,0,0,0,0,0,0,0]
	fmt.Println(cap(slice4))//20
	fmt.Println(len(slice4))//10
      
	//Access Elements of a Slice
	 print("slice2:",slice2[0])
	 
	//appending   slice_name = append(slice_name, element1, element2, ...)
	slice1 = append(slice1, 200,300,400,500)

	//Append One Slice To Another Slice  (  slice3 = append(slice1,slice2...)) {JS SPREAD OPERATOR but pore}
    slice5:= append(slice1,slice4...);

	//from array to slice
	var arr = [...]int{1,2,3,4,5,6};
	var slice3 = arr[0:len(arr)]; // arr[startingIndex: endingIndex(exclusive)] [0:]
	// var slice3 = arr[0:]; // same as above


	//last index value
	lastnumber:= slice1[len(slice1)-1]

	fmt.Println(slice1);
	fmt.Println(slice2);
	fmt.Println(slice3);
	fmt.Println(slice5);
	fmt.Println(len(slice1));
	fmt.Println(lastnumber);



}