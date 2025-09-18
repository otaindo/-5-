package main

import (
 "context"
 "fmt"
 "net/http"
 "sync"
 "time"
)

func handler(w http.ResponseWriter, r *http.Request, wg *sync.WaitGroup) {
 defer wg.Done()
 fmt.Println("Обрабатываем запрос...")
 time.Sleep(2 * time.Second) 
 fmt.Fprintln(w, "Запрос обработан")
 fmt.Println("Запрос завершен")
}

func main() {
 var wg sync.WaitGroup 

 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  wg.Add(1) 
  go handler(w, r, &wg)
 })

 server := &http.Server{
  Addr: ":8080",
 }

 go func() {
  fmt.Println("Сервер запущен на порту 8080...")
  if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
   fmt.Printf("Ошибка запуска сервера: %s\n", err)
  }
 }()

 time.Sleep(1 * time.Second) 
 go http.Get("http://localhost:8080/")
 go http.Get("http://localhost:8080/")

 time.Sleep(3 * time.Second)

 fmt.Println("Остановка сервера...")

 ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
 defer cancel()

 if err := server.Shutdown(ctx); err != nil {
  fmt.Printf("Ошибка при остановке сервера: %s\n", err)
 }
 fmt.Println("Сервер остановлен gracefully")

 wg.Wait() 
 fmt.Println("Все запросы обработаны. Завершение работы.")
}
