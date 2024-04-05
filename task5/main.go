package main

// Разработать программу, которая будет последовательно отправлять значения в канал, 
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

import (
	"flag"
	"fmt"
	"time"
	"math/rand"
)

func Transmitter(ch chan int) {
	ch <- rand.Intn(100)
	fmt.Println("Sent a message")
}

func Receiver(num int) {
	fmt.Println("Received a message: ", num)
}

func main() {
	var liveTime int

	flag.IntVar(&liveTime, "t", 5, "Program live time in seconds")
	flag.IntVar(&liveTime, "time", 5, "Program live time in seconds")
	flag.Parse()

	ch := make(chan int)
	defer close(ch)
	
	timer := time.NewTimer(time.Duration(liveTime) * time.Second)
	
	for alive := true; alive; {
		go Transmitter(ch)
		select {
		case <-timer.C:
			alive = false
			fmt.Println("Close program")
		case mes := <-ch:
			Receiver(mes)
			time.Sleep(1 * time.Second)
		}
	}
}
