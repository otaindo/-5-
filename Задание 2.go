package main

import (
 "fmt"
 "time"
)

func main() {
 done := make(chan struct{})
 go func() {
  for {
   select {
   case <-done:
    return
   default:
    fmt.Println("Привет!")
    time.Sleep(time.Second)
   }
  }
 }()

 time.Sleep(5 * time.Second) 
 close(done)                
}
