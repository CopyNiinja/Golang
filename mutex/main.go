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


	 //In Go, concurrent access to a variable can be unsafe, even if you only read it, if another goroutine might be writing at the same time.
	 //Goroutine 2 could see a partially updated value (because balance += amount is actually multiple CPU instructions: read → add → write).This can cause race conditions.

	Deposit:=func (amount int) {
		mu.Lock()          // lock for writing
		balance += amount
		mu.Unlock()
	}

	Balance:=func () int {
		mu.Lock()          // lock for reading too!
		b := balance
		mu.Unlock()
		return b
	}
		Deposit(1000);
		Balance()

        //Checks his balance hundreds of times per second.Each Balance() call uses mu.Lock(), so all   
		//readers run sequentially.This slows down the system, even though they only read, because readers 
		// don’t need exclusive access—they don’t modify anything.
		//If you want many readers but still exclusive writers, you can use sync.RWMutex:
		//a special kind of lock that allows read-only operations to
		//proceed in parallel with each other, but write operations to have fully exclusive access. 

		/* Time -->
			Reader1 ──RLock──────RUnlock─────
			Reader2 ──RLock──────RUnlock───── (can run in parallel with Reader1)
			Writer  ─────Lock──────Unlock───── (must wait until Reader1 & Reader2 finish)
			Reader3 ──RLock───────────── (waits until Writer finishes)
        */
	var rmu sync.RWMutex
	BalanceWithRWMutex:=func () int {
    rmu.RLock()         // many goroutines can read concurrently
    b := balance
    rmu.RUnlock()
    return b
}

DepositWithRWMutex:=func (amount int) {
    rmu.Lock()          // exclusive lock for writing
    balance += amount
    rmu.Unlock()
}

//RLock() = read lock (multiple readers allowed)
//Lock() = write lock (exclusive)
DepositWithRWMutex(100)
BalanceWithRWMutex();

}