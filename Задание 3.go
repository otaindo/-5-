package main

import "fmt"

func main() {
 ch := make(chan string)

 select {
 case data := <-ch:
  fmt.Printf("Получены данные: %s\n", data)
 default:
  fmt.Println("Данных нет")
 }
}
