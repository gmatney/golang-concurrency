package main

// Where this fun example came from: https://medium.com/@trevor4e/learning-gos-concurrency-through-illustrations-8c4aff603b3
// Has great images of mining gophers

import (
	"fmt"
	"time"
)

func main() {
	mineOre()
	<-time.After(time.Second * 5) // Again, you can ignore this
}

func mineOre() {
	theMine := []string{"rock", "ore", "ore", "rock", "ore", "ore",
		"rock", "ore", "ore", "rock", "ore", "ore", "ore", "ore", "rock"}
	oreChannel := make(chan string, 4)
	minedOreChan := make(chan string, 2)

	go func(mine []string) { // Finder
		for _, item := range mine {
			if item == "ore" {
				oreChannel <- item //send item on oreChannel
			}
		}
	}(theMine)

	go func() { // Ore Breaker
		for foundOre := range oreChannel { //read from oreChannel
			fmt.Println("From Finder: ", foundOre)
			minedOreChan <- "minedOre" //send to minedOreChan
		}
	}()

	go func() { // Smelter
		for minedOre := range minedOreChan { //read from minedOreChan
			fmt.Println("From Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}
	}()

}
