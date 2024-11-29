//Barrier2.go Template Code
//Copyright (C) 2024 Kacper Krakowiak

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Kacper Krakowiak (C00271692)
// Created on 29/11/2024
// Modified by:
// Description:
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// hopefully none
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

// Place a barrier to synchronise each go routine
func doStuff(goNum int, arrived *int, max int, wg *sync.WaitGroup, sharedLock *sync.Mutex, turnstile1 chan bool, turnstile2 chan bool) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	//Phase 1: Arrival at the barrier
	sharedLock.Lock()
	*arrived++
	if *arrived == max { //last to arrive
		//signal all other go routines to pass through through the first turnstile
		for i := 0; i < max-1; i++ {
			turnstile1 <- true
		}
		sharedLock.Unlock()
	} else { //not all here yet we wait until signal
		sharedLock.Unlock()
		<-turnstile1 //wait at the turnstile
	} //end of if-else

	//phase 2: leavong the barrier
	sharedLock.Lock()
	*arrived--
	if *arrived == 0 { //last goroutuine to leave
		// Signal all other goroutines to pass through the second turnstile
		for i := 0; i < max-1; i++ {
			turnstile2 <- true
		}
		sharedLock.Unlock()
	} else { //not all left yet
		sharedLock.Unlock()
		<-turnstile2 //wait @ second
	}
	fmt.Println("Part B", goNum)
	wg.Done()
	return true
} //end-doStuff

func main() {
	totalRoutines := 10
	arrived := 0
	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	var theLock sync.Mutex
	turnstile1 := make(chan bool)
	turnstile2 := make(chan bool)
	for i := 0; i < totalRoutines; i++ {
		go doStuff(i, &arrived, totalRoutines, &wg, &theLock, turnstile1, turnstile2)
	}

	wg.Wait() //wait for everyone to finish before exiting
} //end-main
