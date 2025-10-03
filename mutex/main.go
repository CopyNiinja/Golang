// MutualExclusion:sync.Mutex
// Mutual exclusion (Mutex) ensures that only one goroutine at a time can access
// a critical section (shared resource). like bank balance.. if multiple goroutine change the value together
// there should be inconsistency(race condition).To prevent this->Mutex

package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int = 1000;

func increase(amount int){
	balance+=amount;
}
func main() {

	go increase(300);
	go increase(500);

	time.Sleep(time.Second * 3)// wait  till all the goroutine finished

	fmt.Println(balance) // 1800 but this is not guaranteed.sometimes it can show 1500 or 1300 .cuz the function are running separately (in different goroutine(lightweight thread))

	//mutex
	  var mu sync.Mutex;
	  increaseWithMutex:=func(amount int){
		mu.Lock() // lock 
		defer mu.Unlock() //unlock
		balance+=amount;
	  }
	//Each time a goroutine accesses the variables of the balance, it must call the
	//mutex’s Lock method to acquire an exclusive lock. If some other goroutine has acquired the
	//lock, this operation will block until the other goroutine calls Unlock and the lock becomes
	//available again. The mutex guards the shared variables. 
	fmt.Println(balance)//1800
	 go  increaseWithMutex(400);
	 go increaseWithMutex(500);
	 go increaseWithMutex(40);
      

	 time.Sleep(time.Second * 3)// wait  till all the goroutine finished
	 fmt.Println(balance)//2740

	

}