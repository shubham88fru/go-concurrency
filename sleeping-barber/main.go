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

	// add barbers.
	shop.AddBarber("Barber 1")
	shop.AddBarber("Barber 2")
	shop.AddBarber("Barber 3")
	shop.AddBarber("Barber 4")
	shop.AddBarber("Barber 5")

	// start the barbershop as a goroutine.
	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true

	}()

	//add clients.
	i := 1

	go func() {
		for {
			randomMilliseconds := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return

			case <-time.After(time.Millisecond * time.Duration(randomMilliseconds)):
				shop.addClient(fmt.Sprintf("Client %d", i))
				i += 1

			}
		}
	}()

	// block until the barbershop is closed.

	<-closed

}
