package main

import "fmt"

//a method is just a function with a receiver argument. (JS prototype er moto)
//you can add method to non struct  types too.
type Person struct {
		name string;
		age int;
	}

func (p Person) PrintAge (){
	fmt.Println(p.age);
}

//non struct type

type MyFloat float32;

func (f MyFloat)ToInt()int{
	return int(f)
}
//pointer receiver

type Dimension struct {
	width,height int;
}

func (d *Dimension) Double(){ // if it was not pointer actual value will remain same only the copy/argument 
//                                 //will be updated
	d.height *=2;  // for clear code go syntax: d.height  not (*d).height 
	d.width *=2;
}

func main(){

	var p1=Person{"Faiyaz",26};
    //Method
    p1.PrintAge();

	fmt.Println(p1);
    //other type
	var f1 MyFloat = 123.335;
	fmt.Println(f1.ToInt()) //123
	//pointer receiver
	var d1  Dimension=Dimension{3,4};
	fmt.Println(d1)//{3 4}
	//cuz of pointer it will double the actual value
	d1.Double();
	fmt.Println(d1)//{6 8}



}