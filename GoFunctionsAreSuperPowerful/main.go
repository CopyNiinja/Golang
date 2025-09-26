/*
------Go Functions-------
1. Can return multiple values
2. Named return
3. Can return functions
4. Return slice /maps / array
5. Can return Structs
6.Can define anonymous functions inside main (lambdas/closures)
7.IIFE (Immediately invoked function expression)
8. Error handling
*/

package main

import (
	"errors"
	"fmt"
)

//1. Can return multiple values
func swap (num1 int ,num2 int)(int,int){
	return num2,num1;
}
//2.Named return
func sumAndMultiple(num1 int,num2 int)(sum int,multiple int){

	sum = num1 + num2;
	multiple= num1 * num2;
	return //must put empty return
}
//3.Can return functions
func sumfn(num1 int,num2 int)func(num3 int,num4 int){
  
	return func (num3 int,num4 int){
     println(num1,num2,num3,num4)
	}
}
//4. Return slice /maps / array
func sliceMaker(num1 int,num2 int ,num3 int)[]int{
  return []int{num1,num2,num3}
}

//5. Can return Structs
type Person struct{
	naam string;
	boyosh int;
	bibahito bool;
	gari []string
}
func structMaker (name string,age int ,isMarried bool,cars []string) Person {
	var person Person;
	person.naam=name;
	person.boyosh=age;
	person.bibahito=isMarried;
	person.gari=cars;
	return person;
}

//8. Error Handling
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}


func main(){

	//1. Can return multiple values
    a,b := swap(1,2);
	fmt.Println("SWAP:",a,b) // 2 1

    //2.Named return
	var c,d = sumAndMultiple(10,15);
	fmt.Println("Sum and Multiple:",c,d) //25  150

	//3.Can return functions
	var printer = sumfn(1,2);
	printer(3,4)

	//4. Return slice /maps / array
	var slice1 = sliceMaker(20,24,56);
    fmt.Println("SliceMaker:",slice1) // [20 24 56]

	//5. Can return Structs
	 var manush Person;
	 gariss := []string{"BMW","Pokkhi raj ghora XD"};
	 manush=structMaker("faiyaz",26,false,gariss);
    fmt.Println("Struct: " , manush)

	//6.Can define anonymous functions inside main (lambdas/closures)

	getAge := func(){
		println("Hello DEFINED INSIDE MAIN FUNCTION")
	}
	getAge();

	add:=func(num1 int,num2 int)int{
		return num1+num2;
	}
	fmt.Println("SUM :",add(7,9));

	//7.IIFE (Immediately invoked function expression)
	 result := func(n1 int,n2 int)int{
                      return n1+n2;
	}(4,5);
	fmt.Println("result:",result)

	//8. Error handling
	 value,err := divide(2,6);
	 fmt.Println("VAlue:",value)
	 fmt.Println("Error:",err)




  
	
	














	
 }

