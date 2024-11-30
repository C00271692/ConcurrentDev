// Dining Philosophers deadlock fix through footman Code
// Author: Kacper Krakowiak
// Created: 30/11/24
//GPL Licence

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

// think simulates a philosopher thinking for a random amount of time
func think(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was thinking")
}

// eat simulates a philosopher eating for a random amount of time
func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was eating")
}

// eat simulates a philosopher eating for a random amount of time
// The footman channel ensures that only 4 philosophers can pick up forks at the same time
func getForks(index int, forks map[int]chan bool, footman chan bool) {
	footman <- true            //requesting perm from footman
	forks[index] <- true       //pick up left fork
	forks[(index+1)%5] <- true //pick up right fork
}

// putForks allows a philosopher to put down the forks
// The footman channel is signaled to allow another philosopher to proceed
func putForks(index int, forks map[int]chan bool, footman chan bool) {
	<-forks[index]
	<-forks[(index+1)%5]
	<-footman // signal footman that forks are put down
}

// simulating the life of a philosopher
func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool, footman chan bool) {
	for {
		think(index)
		getForks(index, forks, footman)
		eat(index)
		putForks(index, forks, footman)
	}
	wg.Done() //signals a specific instance of a philosopher is done (not the entire process, its an infinite loop)
}

func main() {
	var wg sync.WaitGroup
	philCount := 5
	wg.Add(philCount) //adding philosophers to the waiting group

	//map to hold the channels for the forks
	forks := make(map[int]chan bool)
	for k := range philCount {
		forks[k] = make(chan bool, 1) //each fork is a buffered channel with 1 capacity
	} //set up forks

	//create a footman channel
	footman := make(chan bool, 4)

	//instantiating the philosophers goroutines
	for N := range philCount {
		go doPhilStuff(N, &wg, forks, footman)
	} //start philosophers
	wg.Wait() //wait here until everyone (10 go routines) is done

} //end of main
