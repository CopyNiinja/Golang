package main

import "fmt"

func main(){
    
	
	for i:=0; i < 5; i++ {
		if i==3{
			continue
		}
		if i==4{
			break
		}
    fmt.Println(i)
  }

      //array iteration
	var arr=[...]int{1,2,3,4,5,6,7,8};
	for i:=0;i<len(arr);i++{
       fmt.Printf("%d\n",arr[i]);
	}
	//range for loop
	for i,val:=range arr{
		fmt.Println(i,"--",val);
	}
    //

}