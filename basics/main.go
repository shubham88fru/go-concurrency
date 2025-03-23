package main

import (
	"fmt"
	"sync"
)

func print(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, x := range words {
		go print(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()
	wg.Add(1)
	print("Hello, World!, again", &wg)
}
