package main

import "fmt"
func main() {
	var marks = 71

	//if else , nested if
	if marks > 80 {
		print("A+")
	} else if marks > 60 {
		if marks > 70 {
			print("A")
		} else {
			print("B+")
		}

	} else {
		print("F")
	}
	//switch

	day := 4

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  default:
    fmt.Println("Not a weekday")
  }

  //multi case
  anotherDay := 5

   switch anotherDay {
   case 1,3,5:
    fmt.Println("Odd weekday")
   case 2,4:
     fmt.Println("Even weekday")
   case 6,7:
    fmt.Println("Weekend")
  default:
    fmt.Println("Invalid day of day number")
  }

}