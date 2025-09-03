package main

import (
	"fmt"
	"math/rand"
)

var MAX_CHICKEN_PRICE int32 = 139545

func main() {
	var chickenChannel = make(chan string)
	var stores = []string{"dmart","reliancefresh","patelmart"}
	for idx := range(stores) {
		go checkPrices(stores[idx],chickenChannel)
	}
	sendMessage(chickenChannel)
}

func checkPrices(store string, chickenChannel chan string) {
	for {
		var newChickenPrice = rand.Int31()
		if newChickenPrice >= MAX_CHICKEN_PRICE {
			chickenChannel <- store
			break
		}
	}
}

func sendMessage(chickenChannel chan string) {
	fmt.Println(<- chickenChannel)
}