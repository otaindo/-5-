import (
 "fmt"
 "time"
)

func main() {

 resultCh := make(chan string)


 go func() {
  time.Sleep(3 * time.Second) 
  resultCh <- "Результат запроса"
 }()


 select {
 case result := <-resultCh:
  fmt.Println("Получен результат:", result)
 case <-time.After(2 * time.Second):
  fmt.Println("Таймаут: сервер не ответил за 2 секунды")
 }
}
