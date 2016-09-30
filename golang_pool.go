package main

import "fmt"
import "os"
import "strconv"
import "sync"

var wg sync.WaitGroup

func producer(replies <-chan int, iterations int) {
	defer wg.Done()
	total := 0
	for index := 1; index <= iterations; index++ {
		reply := <-replies
		total = total + reply
	}
	fmt.Println(total)
}

func consumer(tasks <-chan int, replies chan<- int) {
	defer wg.Done()
	for task := range tasks {
		replies <- task
	}
}

func main() {
	arguments := os.Args[1:]

	consumers, err := strconv.Atoi(arguments[0])
	if err != nil {
		panic(err)
	}

	iterations, err := strconv.Atoi(arguments[1])
	if err != nil {
		panic(err)
	}

	tasks := make(chan int, iterations)
	replies := make(chan int, iterations)

	wg.Add(1)
	go producer(replies, iterations)

	for id := 1; id <= consumers; id++ {
		wg.Add(1)
		go consumer(tasks, replies)
	}

	for index := 1; index <= iterations; index++ {
		tasks <- index
	}
	close(tasks)

	wg.Wait()
}
