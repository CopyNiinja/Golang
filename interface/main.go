package main

import "fmt"

//interface:An interface is like a contract or a set of rules.
//It says: “Any type that has these methods can be treated as this interface.(type)”
/*
Interfaces are satisfied implicitly.

A type can have more methods than the interface asks for → still valid.

But when you use the value as the interface type, you only get access to the interface’s methods.

Type Switch
A type switch can be used to compare the dynamic type of an interface against multiple types.

*/
type Animal interface {

	Speak() string
}

type Cat struct {
	name string
}

func (c Cat) Speak() string {
	return c.name+":MEAAAOOWW";
}
func (c Cat) Run() string {
	return c.name+":fleas away";
}

type Dog struct {
	name string
}

func (d Dog) Speak() string {
	return d.name+":gEHoow GRRGRHHH"
}

type Bird struct {}

func (Bird) Speak() string {
	return "peskdj sdjkopsd"
}

func main() {

	neko:=Cat{
		name: "miyomi",
	}
	pitbull:=Dog{
		name:"pitbully",
	}
	penny := Bird{};

	// i can declare a var as interface type and all the types that has same interface methods can be stored inside this varibale. but it can just use the methods inside interface. not extra method can be applied.like cat has RUN method that is not in the interface. if cat is typed as inteface it cant use the run method;

	Jontu := []Animal{neko,pitbull,penny};
	fmt.Println(Jontu[0].Speak()) //miyomi:MEAAAOOWW

	//SO basically we will use them individually .. if the same code should run for different types. we can make a interface. and put the function parameter as Interface type , use one single function for deferent types arguments. Just they need to have the same method that interface has.

	//type switch
	var f1 float64 = 3.41;
	var i1 int = 73;

	fmt.Printf("%T",ToString(f1))//string
	fmt.Printf("%T",ToString(i1))//string

	//Any type
	PrintAnything(10)//10
	PrintAnything("red")//red
	PrintAnything(true)//true

}
	type Shape interface {
		Area() float64;
	}

	type Circle struct {
		radius float64;
	}
	func (c Circle) Area()float64{
		return 2*3.14* c.radius;
	}
	type Rectangle struct {
		height float64
		width float64
	}
	func (r Rectangle) Area()float64{
		return r.height*r.width
	}

	//USES : Now we can pass any shape(Circle,Rectangle)here 
	func AreaOfTheShape( s Shape)float64{
		return s.Area();
	}




    
	

	
	
