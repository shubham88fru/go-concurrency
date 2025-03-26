package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber += 1
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d!\n", pizzaNumber)
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed += 1
		} else {
			pizzasMade += 1
		}

		total += 1
		fmt.Printf("\nMaking pizza #%d! It will take %d seconds...\n", pizzaNumber, delay)

		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("\n**** We ran out of ingredients for pizza #%d! ****\n", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("\n**** The cook quite while making pizza #%d! ****\n", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("\n**** Pizza #%d is ready! ****\n", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzaMaker *Producer) {
	//keep track of which pizza we are making.
	var i = 0

	//run forever until we receive a quit notification.
	//try to make a pizzas.
	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}

}

func main() {
	//seed random number generator.
	rand.Seed(time.Now().UnixNano())

	//print out a message.
	color.Cyan("Welcome to the pizza shop!\n")
	color.Cyan("--------------------------\n")

	//create a producer.
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//run the producer in the background.
	go pizzeria(pizzaJob)

	//create and run consumer.
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad!")
			}
		} else {
			color.Cyan("Done making pizzas..")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error closing channel!", err)
			}
		}
	}

	//print out the ending message.
	color.Cyan("--------------------------\n")
	color.Cyan("Total pizzas made: %d\n", pizzasMade)
	color.Cyan("Total pizzas failed: %d\n", pizzasFailed)
	color.Cyan("Total: %d\n", total)
	color.Cyan("Goodbye!\n")

	switch {
	case pizzasFailed > 9:
		color.Red("We are going out of business!")
	case pizzasFailed >= 6:
		color.Red("It was a rough day!")
	case pizzasFailed >= 4:
		color.Yellow("We had a few issues today!")
	case pizzasFailed >= 2:
		color.Yellow("We had a couple of issues today!")
	default:
		color.Green("We had a great day!")
	}
}
