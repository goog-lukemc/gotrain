package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"sync"

	"github.com/google/uuid"
)

const (
	LOCALPATH string = "/Users/lukemc/junk"
)

// These are our flag variables
var workers = flag.Int("workers", 1, "Number of Workers")
var testcount = flag.Int("datastrings", 1, "Number of Data Strings")

// Wait group from the sync package
var wg sync.WaitGroup

func main() {
	// Setting a time so we can calculate Duration
	start := time.Now()

	// Don't forget to parse your flags
	flag.Parse()

	// make a channel to safely communicate with our workers
	s := make(chan string)

	// Create a worker pool
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go writeFile(s)
	}

	// create a UUID and set it on the channel
	for x := 0; x < *testcount; x++ {
		ud, err := uuid.NewUUID()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		s <- ud.String()
	}

	// make sure the channel is empty before we close it
	// this helps us support buffered channels
	if len(s) == 0 {
		close(s)
	}

	// waits for all our workers to finish
	wg.Wait()
	f := time.Since(start)
	log.Printf("Workload Count:%v Goroutine Count:%v Duration:%v", *workers, *testcount, f)
}

func writeFile(itemschan chan string) {
	// Range the channel watching as the items come in
	// This loop will continue until items chan is closed
	for data := range itemschan {
		outfile := fmt.Sprintf("%s/%s.txt", LOCALPATH, data)
		err := ioutil.WriteFile(outfile, []byte(data), 0644)
		if err != nil {
			log.Printf("file error%v", err.Error())
		}
	}

	wg.Done()
}

// Let's look at some go routine work patterns
// https://divan.dev/posts/go_concurrency_visualize/

// Exercise 20 mins:
// Write a new program creating a struct that has a map of string string.
// username is the key and password is the value.
// Usernames are u1, u2, ..
// Passwords are p1, p2, ..
// Populate that maps with multiple go routings with 200k
// user password combination in a concurrent safe manner
// Hint read about sync.Mutex and sync.Map

// Bonus: Read this blog on concurency in go
// https://blog.golang.org/pipelines

// Bonus: Leaking memory and go routines
// https://go101.org/article/memory-leaking.html
