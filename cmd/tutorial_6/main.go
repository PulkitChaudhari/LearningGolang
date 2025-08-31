package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var mutex = sync.RWMutex{}
var results = []string{}

func main() {
	var dbIds = []string{"id1","id2","id3","id4","id5"};
	startTime := time.Now()
	for id := range(dbIds) {
		wg.Add(1)
		go dbCall(dbIds[id])
	}
	wg.Wait()
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("Time taken is : ", diff)
	fmt.Println("Results after completion: ", results)
}

func dbCall(dbId string) {
	time.Sleep(time.Millisecond * 2000)
	fmt.Println("DB operation completed for :", dbId)
	save(dbId)
	log()
	wg.Done()
}

func save(dbId string) {
	mutex.Lock()
	results = append(results,dbId)
	mutex.Unlock()
}

func log() {
	mutex.RLock()
	fmt.Println("Current results are :", results)
	mutex.RUnlock()
}