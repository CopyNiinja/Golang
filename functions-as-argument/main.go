package main

//function as argument

import "fmt"

//suppose we want to double or triple each item of an array of number.
//we will pass the function as parameter along with the array . the function will be called with each item of the array.

func doubleOrTriple(numbers *[]int,multiplier func(int)int) []int{
  var modifiedSlice []int;
 for _,num:= range *numbers{
     modifiedSlice= append(modifiedSlice, multiplier(num))
 }
 return modifiedSlice;

}

//function type can be annoying. so a custom type:
type multiplierFn func(int)int

func main() {
	fmt.Println("Function as argument")
	//number slice
	numbers:=[]int{1,2,3,4,5,6,7,8};

	//double multiplier
	doubleMultiplier:=func(num int)int{
       return num *2;
	}
	//Triple multiplier
	tripleMultiplier:=func(num int)int{
       return num *3;
	}
     
	//2x 
	fmt.Println(doubleOrTriple(&numbers,doubleMultiplier)) //[2 4 6 8 10 12 14 16] or another way
	fmt.Println(doubleOrTriple(&numbers,func(i int) int {return i*2})) //[2 4 6 8 10 12 14 16]

	//3x
	fmt.Println(doubleOrTriple(&numbers,tripleMultiplier)) //[3 6 9 12 15 18 21 24]

	//pass by reference. but didnt modify the main array. but its used to stop unnecessary copies.
	fmt.Println(numbers)//[1 2 3 4 5 6 7 8] 
	

	
}