package main

// Реализовать собственную функцию sleep

import (
	"time"
	"fmt"
)

func Sleep(quantity int, dur time.Duration) {
	<-time.After(time.Duration(quantity) * dur)
}

func main() {
	start := time.Now()
	Sleep(5, time.Second)

	fmt.Printf("Hello - %2.0f сек", time.Since(start).Seconds())
}
