package models

import (
	"fmt"
	"net/http"
)

// Service
type Service struct {
	Timestamp   string
	Name        string
	URL         string
	Healthy     bool
	Description string
}

// NewService is constructor for service.
func NewService(name string, url string) Service {

	service := Service{Name: name, URL: url}

	return service
}

func (s Service) CheckHealth() bool {

	// implement the logic to check http get requst and return if 200 healthy or not

	resp, err := http.Get(s.URL)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	fmt.Println("resp.StatusCode: ", resp.StatusCode)
	// return resp.StatusCode == 200
	return resp.StatusCode == 200
}
