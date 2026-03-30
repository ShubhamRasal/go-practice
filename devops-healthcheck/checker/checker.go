package checker


import (

  "fmt"
  "github.com/fatih/color"
  "github.com/ShubhamRasal/go-practice/devops-healthcheck/models"

)

func PrintStatus(s models.Service){
  status := "Healthy"
  
  if s.Healthy == false {
    status = "Unhealthy"
    
    msg :=  fmt.Sprintf("Name: %s | Port: %d  | %s ", s.Name, s.Port, status)
    color.Red(msg)

   }else{
    msg :=  fmt.Sprintf("Name: %s | Port: %d  | %s ", s.Name, s.Port, status)
    color.HiGreen(msg)
   }

}
