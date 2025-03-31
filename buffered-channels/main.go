package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		//got a meessage.
		i := <-ch
		fmt.Println("Got", i, "from channel")

		//simulate doing a lot of work
		time.Sleep(time.Second * 1)
	}
}

func main() {
	ch := make(chan int, 10) //buffered channel - channel with a size. Default is 0

	go listenToChan(ch)

	for i := 0; i < 100; i++ {
		fmt.Println("Sending ", i, "to channel..")
		ch <- i
		fmt.Println("Sent ", i, "to channel")
	}

	fmt.Println("All done.")
	close(ch)
}
