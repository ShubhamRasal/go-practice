package main


import (
  "fmt"
)

func main(){
  fmt.Println("hello, loop")

  // classic loop
  for i:= 0; i<=5 ; i++ {
    fmt.Println(i) 
  }
 
 // range - iterate array/slice
 var servers []string
 servers = []string{"server-01", "server-02", "server-03"}
  
 fmt.Println("using classic loop:")
 // classic
 for i := 0; i < len(servers); i++ {
   fmt.Println("index ", i, "server:", servers[i])
 }


 fmt.Println("using range:")
 // for index, index th value, 
 for i, s := range servers {
	fmt.Printf("[%d] => %s \n", i, s)
 }
 
 fmt.Println("skip index with _")
 for _, s := range servers {
      fmt.Println(s)
 }

  fmt.Println("using for in while-style")
 // while
 retries := 0
 for retries < 5 {

   if retries == 3{
    
     retries++
     continue
   } 
   
   fmt.Println(retries)

   retries++

  }
}























