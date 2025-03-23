package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	defer m.Unlock()
	msg = s
}

func main() {
	msg = "Hello, World!"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Goodbye, World!", &mutex)
	go updateMessage("Goodbye, Cruel World!", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
