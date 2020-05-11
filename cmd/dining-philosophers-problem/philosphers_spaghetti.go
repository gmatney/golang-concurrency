package main

/*
 Examples:
 https://en.wikipedia.org/wiki/Dining_philosophers_problem
 https://play.golang.org/p/rXCotNNY24
 https://www.golangprograms.com/illustration-of-the-dining-philosophers-problem-in-golang.html
 https://github.com/will-zegers/Go-DiningPhilsophers/tree/master/forks

*/

import (
	"fmt"
	"hash/fnv"
	"log"
	"math/rand"
	"sync"
	"time"
)

var ph = []string{"Aristotle", "Plato", "Kant", "Confucius", "Buddha", "Nietzche"}

const think = time.Second / 100 // Mean think time
const eat = time.Second / 100   // Mean eat time

func randomSleep(phName string, meanTime time.Duration) { // Use different seeds for random, their names
	h := fnv.New64a()
	h.Write([]byte(phName))
	rg := rand.New(rand.NewSource(int64(h.Sum64())))
	rSleep := func(t time.Duration) {
		time.Sleep(t/2 + time.Duration(rg.Int63n(int64(t))))
	}
	rSleep(eat)
}

func diningProblem(phName string, leftHand, rightHand *sync.Mutex) {
	for {
		//Technically everyone could grab the left fork at exactly the same time.
		//Technically I could roll ten nat 20s in a row in DnD too.
		//Will do another example that doesn't have this possibility.
		log.Println(phName, "Hungry")
		leftHand.Lock()  // Keep thinking until left is available, when available pick up.
		rightHand.Lock() // Keep thinking until right is available, when available pick up.
		log.Println(phName, "Eating")
		randomSleep(phName, eat)
		leftHand.Unlock() // put down forks
		rightHand.Unlock()
		log.Println(phName, "Thinking")
		randomSleep(phName, think)
	}
}

func main() {
	forever := sync.WaitGroup{}
	forever.Add(1)
	fork0 := &sync.Mutex{}
	forkLeft := fork0
	for i := 1; i < len(ph); i++ {
		forkRight := &sync.Mutex{}
		go diningProblem(ph[i], forkLeft, forkRight)
		forkLeft = forkRight
	}
	forever.Wait() //Philospher's dine forever.
	fmt.Println("Table empty")
}
