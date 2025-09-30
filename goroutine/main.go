//A goroutine is a lightweight thread managed by the Go runtime.
//They’re so lightweight that you can easily create thousands of them without blowing up your memory usage.
//To start a goroutine, you simply prefix a function call with the go keyword.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
   
	 greet() 
	 go slowGreet() //executed as a goroutine.
	 greet() 

	 /*If you check console , it only shows two greet() functions returns
	 	The reason your slowGreet() doesn’t show up is because your main function finishes before the goroutine has time to run.
		go slowGreet() starts a new goroutine, but it doesn’t block.It schedules slowGreet() to run concurrently.

		*/

		//to wait for the goroutine to complete before exiting: 
		//way 1: keep the main awake for some seconds(but you need to know how much time will take for goroutines ) 
		time.Sleep(5*time.Second) //it will keep awake the main function till the goroutines finish their works

		//way 2: sync.WaitGroup

		var wg sync.WaitGroup

	    greet()

		wg.Add(1) //tell wait group that we have 1 goroutine

		//instead of calling the slow function , make an IIFE f  
		go func(){
			defer wg.Done();  //defer : it will execute in lIFO   wg.Done(): mark goroutine as done
			slowGreet();
		}() 
        //to know the tread is not blocked ny slowGreet
		greet()

		wg.Wait(); // wait till goroutines to finish (till wg.Done() is not called)





}

func greet(){
	fmt.Println("Hello")
}
func slowGreet(){
	time.Sleep(3*time.Second); //replicates slow task 
	fmt.Println("H  E  L  L O (slowly)")
}