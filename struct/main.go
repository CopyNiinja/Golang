package main

import (
	"errors"
	"fmt"
	"time"
)

//Outside main {First letter capital will make it available in different packages}
type User struct {
	name string;
	age int;
}

type Admin struct {
	firstName string
	lastName string
	 bio string
	 age int
	 createdAt time.Time 
}

func newAdmin(fname,lname,bio string,age int)(*Admin,error){
	if(fname=="" && lname==""){
     return nil,errors.New("Firstname and lastname are required!!")
	}
 return &Admin{firstName: fname,lastName: lname,bio: bio,age: age,createdAt: time.Now()},nil
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

	//
	var admin1 = Admin{ firstName: "Faiyaz",lastName:"Ahmed",bio: "dsd",age: 26,createdAt: time.Now()}
   var admin2  ,err= newAdmin("Mushfiqur","Rahim","hehehe",43);
	fmt.Println(admin1)
	fmt.Println(*admin2,err);
}