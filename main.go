package main

import "fmt"


type Mutex struct {
  Count  int           
  signal chan struct{} 
}


func (m *Mutex) Notify() {
  m.signal <- struct{}{}
}


func (m *Mutex) Wait() {
  for i := 0; i < m.Count; i++ {
    <-m.signal 
  }
}

func main() {
  const goroutinesCount = 7

 
  m := Mutex{
    Count:  goroutinesCount,
    signal: make(chan struct{}, goroutinesCount),
  }

  for i := 0; i < goroutinesCount; i++ {
    go func(index int) { 
      defer m.Notify() 
      fmt.Printf("Горутина %d: Привет, мир!\n", index)
    }(i)
  }

  // Ожидаем завершения всех горутин
  m.Wait()
}
