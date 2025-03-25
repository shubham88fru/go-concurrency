package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumerOfPizzas = 10

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
	if pizzaNumber <= NumerOfPizzas {
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
		fmt.Printf("Making pizza #%d! It will take %d seconds...", pizzaNumber, delay)

		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("**** We ran out of ingredients for pizza #%d! ****", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("**** The cook quite while making pizza #%d! ****", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("**** Pizza #%d is ready! ****", pizzaNumber)
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
	// var i = 0

	//run forever until we receive a quit notification.
	//try to make a pizzas.
	for {
		// currentPizza := makePizza(i)
		//try to make a pizza.
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

	//print out the ending message.
}
