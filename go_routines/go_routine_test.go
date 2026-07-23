package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println(phrase)
	doneChan <- true

}
func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello1 :", phrase)
	doneChan <- true
	close(doneChan)
}

func TestRoutines(t *testing.T) {
	dones := make([]chan bool, 4)
	// done := make(chan bool)
	dones[0] = make(chan bool)
	go slowGreet("How .. are .. you", dones[0])
	dones[1] = make(chan bool)
	go greet("I hope u are fine", dones[1])
	dones[2] = make(chan bool)
	go greet("ya just a little headache", dones[2])
	dones[3] = make(chan bool)
	go greet("Get well soon", dones[3])

	for _, done := range dones {
		<-done

	}
}

func TestRoutinesTwo(t *testing.T) {
	done := make(chan bool)
	go slowGreet("How .. are .. you", done)
	go greet("I hope u are fine", done)
	go greet("ya just a little headache", done)
	go greet("Get well soon", done)

	for range done {
		fmt.Println(done)
	}

}
