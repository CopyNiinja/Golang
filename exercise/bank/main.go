package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const BalanceFileName= "balance.txt"

func ReadBalance()(float64,error){
  data ,err:=  os.ReadFile(BalanceFileName);
    
  if(err != nil){
  return 0,errors.New("Failed to read balance!")
  }
  balance,err := strconv.ParseFloat(string(data),64);

  return balance,err
}

func WriteBalance(amount ,balance float64, t string)(float64,error){
	//withdraw
	if t=="withdraw"{
		if amount > balance{
			return balance, errors.New("You dont have enough balance!");
		}else{
			sliceOfByte:= fmt.Sprint(balance-amount);
		 errorInWritingFile:=	os.WriteFile(BalanceFileName,[]byte(sliceOfByte),0664)
		 if errorInWritingFile!=nil{
			return balance, errors.New("There was a problem!");
		 }else {
			return balance+amount,nil;
		 }
		}
      //deposit
	}else{
		totalAmount := strconv.FormatFloat(amount+balance,'f', 4, 64)
        errorInWritingFile:= os.WriteFile(BalanceFileName,[]byte(totalAmount),0664)
		if errorInWritingFile!=nil {
			return balance,errors.New("There was a problem!")
		}else {
			return balance+amount,nil
		}
		
	}


}
func main(){


	fmt.Println("Welcome to the bank");
	

	//infinity loop

	for {
		bankBalance ,err := ReadBalance();

	if(err != nil){
		fmt.Println(err);
		return
	}

		var choice int;
		var depositMoney float64;
		var withdrawMoney float64;
		fmt.Println("How can I help you? Choose one of the four options->")
		fmt.Println("1.Check you bank balance")
		fmt.Println("2.Deposit money to you account")
		fmt.Println("3.Withdraw money from your account")
		fmt.Println("4.Exit")
		fmt.Print("You choose: ")
		fmt.Scan(&choice);

		switch choice {
		case 1:
			fmt.Printf("Your bank balance is :%.4f\n",bankBalance)
		case 2:
			fmt.Print("How much money you want to deposit:")
			fmt.Scan(&depositMoney);
			 result,err:= WriteBalance(depositMoney,bankBalance,"deposit")
			if(err !=nil){
                  fmt.Println("FAILED:",err);
				  return ;
			 }else{
               fmt.Println("Your current balance is: ",result);
			 }	
		case 3:
			fmt.Print("How much money you want to withdraw:")
			fmt.Scan(&withdrawMoney);
			 result,err:= WriteBalance(withdrawMoney,bankBalance,"withdraw")
			 if(err !=nil){
                  fmt.Println("FAILED:",err);
				  if err.Error() =="You dont have enough balance!"{
					continue;
				  }
				  return ;
			 }else{
               fmt.Println("Your current balance is: ",result);
			 }
			
	    default:
			fmt.Println("-----------------------")
			fmt.Println("Thank you for using our bank.")
			return

		
		
		}
	}

	
}