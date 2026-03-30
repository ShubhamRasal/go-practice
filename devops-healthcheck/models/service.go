package models


// Service 
type Service struct {
  Name string
  Port int
  Healthy bool
  desc string
}

// NewService is constructor for service.
func NewService(name string, port int) Service {

   // validport

   isHealthy := ValidatePort(port)

  service :=  Service{Name: name,Port: port, Healthy: isHealthy}

  return service
}


func ValidatePort(port int) bool {

	if port <= 0 || port >= 65535 {
		return false

	} else {
	    return true

	}
}
