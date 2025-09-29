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

}
//Any Type (like typescript)
func PrintAnything(val any){
 fmt.Println(val)
}