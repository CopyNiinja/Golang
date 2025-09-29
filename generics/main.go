package main

import "fmt"

//Generics (Typescript :D)



func add[T int|float64,K bool](  num1,num2 T,doPrint K)T{
	total := num1+num2
	if doPrint {
		fmt.Println(total)
	}
 return total
}

func main() {


	//add two int
	result:= add(10,44,true);//54    //result type is inferred automatically from the function arguments
    
	
	//add two float
	fmt.Println(result,add(3.41,6,false))//  54  9.41
}