// function that can have any number of parameter

package main

import "fmt"

func printer(num1,num2 int,isBool bool,nums ...int){   //must be last parameter
  fmt.Println(num1,"-",num2,"-",isBool,"-",nums); //1 - 2 - false - [45 67 45 88 55 66]
} 

func main(){
  printer(1,2,false,45,67,45,88,55,66)

}