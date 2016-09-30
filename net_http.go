package main

import "bytes"
import "fmt"
import "encoding/json"
import "gopkg.in/cheggaaa/pb.v1"
import "net/http"
import "os"
import "strconv"
import "sync"
import "time"

var wg sync.WaitGroup

type Response struct {
	status_code int
	size        int
}

func producer(responses <-chan Response, iterations int) {
	defer wg.Done()

	var status_codes = make(map[int]int)

	size := 0
	progress_bar := pb.StartNew(iterations)
	progress_bar.SetRefreshRate(time.Millisecond * 100)
	for index := 1; index <= iterations; index++ {
		response := <-responses
		if _, present := status_codes[response.status_code]; !present {
			status_codes[response.status_code] = 0
		}
		status_codes[response.status_code]++
		size = size + response.size
		progress_bar.Increment()
	}
	bytes, err := json.MarshalIndent(status_codes, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

	fmt.Println(size)
}

func consumer(requests <-chan int, responses chan<- Response) {
	defer wg.Done()

	for _ = range requests {
		timeout := time.Duration(5 * time.Second)
		client := http.Client{
			Timeout: timeout,
		}
		response, err := client.Get("https://bitbucket.org")
		if err != nil {
			responses <- Response{status_code: 999, size: 0}
			return
		}

		defer response.Body.Close()

		buffer := &bytes.Buffer{}
		err = response.Write(buffer)
		if err != nil {
			responses <- Response{status_code: response.StatusCode, size: 0}
			return
		}

		size := buffer.Len()
		responses <- Response{status_code: response.StatusCode, size: size}
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

	requests := make(chan int, iterations)
	responses := make(chan Response, iterations)

	wg.Add(1)
	go producer(responses, iterations)

	for index := 1; index <= consumers; index++ {
		wg.Add(1)
		go consumer(requests, responses)
	}

	for index := 1; index <= iterations; index++ {
		requests <- index
	}
	close(requests)

	wg.Wait()
}
