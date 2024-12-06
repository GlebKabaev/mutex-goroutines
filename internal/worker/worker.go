package worker

import (
	"fmt"
	"github.com/GlebKabaev/mutex-goroutines/pkg/mutex"
)

func RunWorkers(count int) {
	m := mutex.New(count)

	for i := 0; i < count; i++ {
		go func(id int) {
			defer m.Unlock()
			fmt.Printf("Горутина %d: Привет, мир!\n", id)
		}(i)
	}

	m.Wait()
}
