package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func checkService(name, url string, wg *sync.WaitGroup) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("name: ", name, "url: ", url, "Status:", resp.Status)

	resp.Body.Close()

	wg.Done() // wg.Sub(1)
}

func main() {
	fmt.Println("hello world")

	var wg sync.WaitGroup // GPS, tracker, counter

	startTime := time.Now()

	wg.Add(1)
	go checkService("api1", "https://httpbin.org/status/200", &wg)

	wg.Add(1)
	go checkService("api2", "https://httpbin.org/status/200", &wg)

	wg.Add(1)
	go checkService("api3", "https://httpbin.org/status/200", &wg)

	wg.Add(1)
	go checkService("api4", "https://httpbin.org/status/200", &wg)

	wg.Add(1)
	go checkService("api5", "https://httpbin.org/status/200", &wg)

	wg.Add(1)
	go checkService("api6", "https://httpbin.org/status/200", &wg)

	wg.Add(1)
	go checkService("api7", "https://httpbin.org/status/200", &wg)

	wg.Add(1)
	go checkService("api8", "https://httpbin.org/status/200", &wg)

	wg.Add(1)
	go checkService("api9", "https://httpbin.org/status/200", &wg)

	wg.Wait()

	fmt.Println("total time :", time.Since(startTime))
}
