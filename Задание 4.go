package main

import (
 "fmt"
 "sync"
)

type Task struct {
 Priority int 
 Data     string
}

func main() {
 highPriority := make(chan Task, 10) 
 lowPriority := make(chan Task, 10)  
 done := make(chan bool)             

 var wg sync.WaitGroup

 go func() {
  for task := range highPriority {
   wg.Add(1)
   go func(t Task) {
    defer wg.Done()
    fmt.Printf("Обрабатываем (высокий приоритет): %s\n", t.Data)
   }(task) 
  }
  done <- true 
 }()

 go func() {
  for task := range lowPriority {
   wg.Add(1)
   go func(t Task) {
    defer wg.Done()
    fmt.Printf("Обрабатываем (низкий приоритет): %s\n", t.Data)
   }(task) 
  }
  done <- true 
 }()

 highPriority <- Task{Priority: 0, Data: "Срочный заказ"}
 lowPriority <- Task{Priority: 1, Data: "Обычный заказ"}
 highPriority <- Task{Priority: 0, Data: "Важное обновление"}

 close(highPriority)
 close(lowPriority)

 <-done
 <-done  

 wg.Wait() 
 fmt.Println("Все задачи обработаны.")
}
