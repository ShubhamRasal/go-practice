package main


import (
	"fmt"
	"github.com/ShubhamRasal/go-practice/devops-healthcheck/models"
	"github.com/ShubhamRasal/go-practice/devops-healthcheck/checker"
)

func main() {

	fmt.Println("welcome to devops-healthcheck")
	
	services := []models.Service{
   	   {Name: "gateway",  Port: 8080  , Healthy: true },
    	   {Name: "postgres",  Port: 5432 , Healthy: false },
    	   {Name: "frontend",  Port:  443 , Healthy:  true },
       }

	for _, svc  := range services {
	
		checker.PrintStatus(svc)
	}
}
