// func FunctionName(param1 type, param2 type) type {
//   // code to be executed
//   return output
// }

package main

import "fmt"

func main(){
	hello("Faiyaz");
	 var arr = [...]int{1,2,3,4}; // array : Pass by value where slice is pass by ref
	 printArray(arr,7);
	 print("----",arr[1]);//2 (if it was slice, it will be changed t0 100 by the function)

	 print(sumOfTheNumbers(1,2,3))//6


}

func sumOfTheNumbers(num1 int,num2 int,num3 int)int{

	return num1+num2+num3;
}

func printArray(array [4]int,num int){
	array[1]=100;
	for i:=0;i<len(array);i++{
		fmt.Println(i,"-",array[i])
	}
	print(num)
	
}

func hello(name string){
	fmt.Println("HEllo :",name)
}
