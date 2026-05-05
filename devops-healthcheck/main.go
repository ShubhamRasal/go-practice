package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ShubhamRasal/go-practice/devops-healthcheck/models"
)

// db array of services
var services map[string]models.Service

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// validate the request if it post or not
	// wrong input early return
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// accept the request body, read it and parse it into models.Service
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var service models.Service
	err = json.Unmarshal(body, &service)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusInternalServerError)
		return
	}

	service.Healthy = service.CheckHealth()
	now := time.Now()
	// format it into yyyy-mm-dd hh:mm:ss
	formattedTime := now.Format("2006-01-02 15:04:05")
	service.Timestamp = formattedTime
	services[service.Name] = service

	// return the service details in json format
	// todo: explain later
	json.NewEncoder(w).Encode(service)
}

func getAllServicesHandler(w http.ResponseWriter, r *http.Request) {

	// validate the request if it is get or not
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("total services: ", len(services))

	json.NewEncoder(w).Encode(services)
}

func runAll(w http.ResponseWriter, r *http.Request) {
	// validate the request if it is post or not
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	for key, value := range services {
		value.Healthy = value.CheckHealth()
		now := time.Now()
		// format it into yyyy-mm-dd hh:mm:ss
		formattedTime := now.Format("2006-01-02 15:04:05")
		value.Timestamp = formattedTime
		// update db map with new healthy value
		services[key] = value
	}
	json.NewEncoder(w).Encode(services)
}

func main() {

	fmt.Println("hello http-server")

	services = make(map[string]models.Service)

	// registers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/healthcheck", healthcheckHandler)
	http.HandleFunc("/services", getAllServicesHandler)
	http.HandleFunc("/runall", runAll)

	fmt.Println("Server running on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
