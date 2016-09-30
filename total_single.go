package main

import "fmt"
import "os"
import "strconv"

func main() {
	arguments := os.Args[1:]

	iterations, err := strconv.Atoi(arguments[0])
	if err != nil {
		panic(err)
	}

	total := 0
	for number := 1; number <= iterations; number++ {
		total = total + number
	}

	fmt.Println(total)
}
