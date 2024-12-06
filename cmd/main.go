package main

import (
	"github.com/GlebKabaev/mutex-goroutines/internal/worker"
)

func main() {
	worker.RunWorkers(3)
}
