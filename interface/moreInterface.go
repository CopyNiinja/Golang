package main

import (
	"fmt"
	"strconv"
)

func ToString(val interface{})interface{} {
	
	
	switch v:= val.(type){ // concrete type
	case string:
		return v
    case int:
		return strconv.Itoa(v);
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64) 
	default:
		return fmt.Sprintf("%v", v) 	

	}
	//another way

	value,ok :=val.(int); //if the value is int type then ok will be true,value will be the number argument
	if(ok) {//means the value is int type
		value+=1; //value type int not interface{} anymore
        
		} else{
			//means the value is not int
			//value will be 0 (default value)
		}
		return val

}
//Any Type (like typescript)
func PrintAnything(val any){
 fmt.Println(val)
}