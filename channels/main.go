// goroutine runs independently. You don’t “call” it, you just fire it off.
// a goroutine itself does not return a value directly.
// channel: works as a communication mechanisms.Channels are a typed conduit
// through which you can send and receive values with the channel operator, <-
// You can also pass a pointer or shared variable, but channels are idiomatic Go.
package main

import (
	"fmt"
	"time"
)

func main() {
  //to make a channel (channel can be any type)
  ch1 :=make(chan string);
   
  greet(1);

  //starts a goroutine
  go func(){
    slowGreet(1);
	ch1<-"DONE";
  }()

  greet(2);

  var tempString string;

  greet(3);
   
  tempString = <-ch1; //holding the value inside a temp 
   
  greet(4);

  fmt.Println(tempString);

  greet(5);

  //------------------Prints-------------------
  /*
  	Hello 1
	Hello 2
	Hello 3 ( then wait few seconds)
	H E L L O :  1  ...
	Hello 4
	DONE
	Hello 5
   -------------
    code flow:
		greet(1)// prints "Hello 1"
				go func() { // goroutine starts
				slowGreet(1) // sleeps 3s, then prints
				ch1 <- "DONE"
				}()
		greet(2)// prints "Hello 2"
		greet(3)// prints "Hello 3"

							// ---- THIS LINE ----
							tempString = <-ch1 // blocks until goroutine sends
							// -------------------

							greet(4)    // prints "Hello 4"
							fmt.Println(tempString) // prints "DONE"
							greet(5)    // prints "Hello 5"`
  */

	

}

func greet(num int){fmt.Println("Hello",num)}
//replicate slow tasks that takes longer time
func slowGreet(num int){ time.Sleep(3*time.Second) ; fmt.Println("H E L L O : ",num," ...")}