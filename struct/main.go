package main

import "fmt"

//Outside main {First letter capital will make it available in different packages}
type User struct {
	name string;
	age int;
}

func main(){
    
	//Inside main
	type Person struct {
		name string;
		age int;
	}

	var parson1 Person;
	parson1.age=26;
	parson1.name = "faiyaz";

	//outside
	var user1 User;
	user1.age=40;
	user1.name="Rahim";


	fmt.Println(parson1)
	fmt.Println(user1)
}