package main

import (
	"fmt"
	"slices"
)

//needed for v:1.21+ methods only
func main(){
	//Modifying the slice affects the original array!!! (PASS BY REFERENCE)
	var arr = [10]int{1,2,3,4,5}
    var slc = []int{10,11,12};

	slcFromArr := arr[0:];
   //----before updating-----
	fmt.Println(arr)//
	//----updating slice-----
	slcFromArr[0]=99;
   //----after updating-----
    fmt.Println(arr)



	fmt.Println(slc)
	fmt.Println(slcFromArr)

//--------------new v:1.21+ ---------------------------------
	s := []int{3, 1, 4, 1, 5}

    fmt.Println("Original:", s)

    // clone
    t := slices.Clone(s)
    fmt.Println("Cloned:", t)

    // contains
    fmt.Println("Contains 4?", slices.Contains(s, 4))

    // sort
    slices.Sort(s)
    fmt.Println("Sorted:", s)

    // reverse
    slices.Reverse(s)
    fmt.Println("Reversed:", s)

    // delete elements 1 to 3
    s = slices.Delete(s, 1, 3)
    fmt.Println("After delete:", s)

    // insert at index 1
    s = slices.Insert(s, 1, 9, 8)
    fmt.Println("After insert:", s)

   

}
