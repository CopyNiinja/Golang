package greet

import (
	"fmt"
)



func SayHello(name string)  {
     ping()
	fmt.Println("Hello,",name)

  }

func ping(){
	fmt.Println("DING DONG, but its private!(first word small=>p )")
}
