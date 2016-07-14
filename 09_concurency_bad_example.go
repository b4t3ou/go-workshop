package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var items []int

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go appendToItems(3, &wg, "one")
	go appendToItems(2, &wg, "two")

	wg.Wait()
	fmt.Println(items)
	fmt.Println("Done")
}

func appendToItems(multiplier int, mainWG *sync.WaitGroup, name string) {
	wg := sync.WaitGroup{}
	counter := random(1, 100)
	wg.Add(counter - 1)

	for i := 1; i < counter; i++ {
		go func(counter int) {
			duration := random(1, 100)
			time.Sleep(time.Duration(random(1, 100)) * time.Millisecond)
			items = append(items, duration)
			fmt.Println(name + " - " + strconv.Itoa(duration))
			wg.Done()
		}(i)
	}

	wg.Wait()
	mainWG.Done()
}

func random(min, max int) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(max-min) + min
}