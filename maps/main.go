package main

import "fmt"

// var map1 = map[KeyType]ValueType{ key1:value1,key2:value2,...}

func main(){


	//EMPTY MAP  = nil (SYNTAX IS WEIRD I KNOW)
	 var map1 map[string]int;
	 //EMPTY MAP  using make !== nil
	 var map2 = make(map[int]string);
	
	fmt.Println(map1==nil)//true if its nil than no insertion deletion
	fmt.Println(map2==nil)//false


	// MAp
	map3 := map[string]int{ "banana":1,"apple":4,"orange":3};
    map3["guava"]=10;
	fmt.Println(map3["guava"])

	//Map always return two value  ( value and exists)

	value, exists := map3["pineapple"];
	var value2, exists2 = map3["apple"];
	fmt.Println(value,exists)//0 false    0->cuz it was int type value,and default=0
	fmt.Println(value2,exists2)//4 true   

	//Remove Element from Map
	delete(map3,"apple")
	fmt.Println(map3);

	//Iterate Over Maps
	for i,val:= range map3{
       fmt.Println(i,val);
	}
}