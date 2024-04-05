package main

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные
// данные из канала и выводят в stdout. Необходима возможность
// выбора количества воркеров при старте.

import (
	"flag"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"os"
	"os/signal"
	"syscall"
)

func Worker(i int, ch chan interface{}) {
	for {
		fmt.Printf("Worker %d received data: %v\n", i, <-ch)
	}
}

func main() {
	var numWorkers int
	flag.IntVar(&numWorkers, "w", 3, "Number of workers to run concurrently")
	flag.IntVar(&numWorkers, "workers", 3, "Number of workers to run concurrently")
	flag.Parse()

	ch := make(chan interface{}, numWorkers)
	defer close(ch)

	for i := 0; i < numWorkers; i++ {
		go Worker(i + 1, ch)
	}
	// генератор случайных данных
	go func () {
		for {
			ch <- gofakeit.Sentence(5)
			ch <- gofakeit.Name()
			ch <- gofakeit.Int64()
		}
	}()

	sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, os.Interrupt, syscall.SIGINT)
    <-sigChan
}
