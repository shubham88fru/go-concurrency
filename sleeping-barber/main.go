package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// variables.
var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	// seed a random number generator.
	rand.Seed(time.Now().UnixNano())

	// welcome message
	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	// create channels if needed.
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create a ds to represent the barber shop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	color.Green("The shop is open for the day!")
	fmt.Println(shop)

	// add barbers.

	// start the barbershop as a goroutine.

	// block until the barbershop is closed.

}
