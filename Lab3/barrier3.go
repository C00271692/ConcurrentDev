//Barrier.go Template Code
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
// Created on 30/9/2024
// Modified by:
// Description:
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// many
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, arrived *int, max int, wg *sync.WaitGroup, sharedLock *sync.Mutex, theChan chan bool) bool {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)

	//we wait here until everyone has completed part A
	sharedLock.Lock()
	*arrived++
	if *arrived == max { //last to arrive -signal others to go
		for i := 0; i < max-1; i++ {
			theChan <- true
		}
		*arrived = 0
		sharedLock.Unlock()
	} else { //not all here yet we wait until signal
		sharedLock.Unlock()
		<-theChan
	} //end of if-else
	fmt.Println("Part B", goNum)
	sharedLock.Lock()
	*arrived++

	if *arrived == max {
		for i := 0; i < max-1; i++ {
			theChan <- true
		}
		*arrived = 0
		sharedLock.Unlock()
	} else {
		sharedLock.Unlock()
		<-theChan
	}

	fmt.Println("PartC", goNum)
	return true
} //end-doStuff

func main() {
	totalRoutines := 10
	arrived := 0
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	var theLock sync.Mutex
	theChan := make(chan bool) //use unbuffered channel in place of semaphore

	for i := 0; i < totalRoutines; i++ { //create go routines here
		go doStuff(i, &arrived, totalRoutines, &wg, &theLock, theChan)
	}
	wg.Wait() //wait for everyone to finish before exiting
} //end-main
