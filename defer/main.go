package main

import "fmt"

//defer use a stack (LIFO) .
//defer waits till the function ends
// WHen ypu open a folder immediately you can defer fs.close()  it will execute after the function end!!

func main(){

	fmt.Println("Starting line")

	 defer fmt.Println("defer line 1")
	 fmt.Println("Middle line");
	 defer fmt.Println("defer line 2")
	fmt.Println("End line")

	/*------ print ------- 

		Starting line
		Middle line
		End line
		defer line 2
		defer line 1
	*/
}