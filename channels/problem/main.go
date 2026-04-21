package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkService(name, url string) string {

	resp, err := http.Get(url)
	if err != nil {
		return "failed"
	}

	// fmt.Println("name: ", name, "url: ", url, "Status:", resp.Status)

	resp.Body.Close()

	return fmt.Sprintf("%s : %s -- %d ", name, url, resp.Status)
}

func main() {
	fmt.Println("welcome to channels session")

	go checkService("api", "https://example.com")

	time.Sleep(3 * time.Second)

}
