package main

//UNLIKE C,C++ NO POINTER ARITHMETIC IN GO :(
import "fmt"
func main(){

	var x int = 10;
	fmt.Println("x->",x); //10
	fmt.Println("address of x->",&x); //0xc00008c060
	fmt.Println("value of the address of x->",*(&x));//10

   //initialize a pointer
	var ptr *int;
	ptr = &x;

	fmt.Println("pointer holding address ptr->",ptr);//0xc00008c060
	fmt.Println(" the value is located in address ptr->",*ptr); //10

	//pointer can be used to pass the reference

	a:=23;

	change:=func(numPointer *int){
		*numPointer= 500;
	}
	fmt.Println("a:",a);//23
	change(&a);
	fmt.Println("a:",a);//500

	//If you pass a big struct to a function, Go copies the whole thing.
    //Using a pointer avoids that cost.a



type Person struct {
    Name string
    Age  int
    Data [1000000]int // huge array so better just pass the address !!!
}

//   func updateAge(p *Person, newAge int) {
//     p.Age = newAge
// }

	


}