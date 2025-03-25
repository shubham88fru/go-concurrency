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

	example2()
}

var wg2 sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func example2() {
	//variable for bank balance.
	var bankBalance int
	var balance sync.Mutex

	//print out starting values.
	fmt.Printf("Initial account balance: $%d.00", bankBalance)

	//define weekly revenue.
	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts ", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}
	wg2.Add(len(incomes))

	//loop through 52 weeks and print how much is made; keep a running total.
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg2.Done()

			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("Week %d: %s: $%d.00\n", week, income.Source, income.Amount)
			}
		}(i, income)
	}
	wg2.Wait()

	//print out final balance.
	fmt.Printf("Final account balance: $%d.00", bankBalance)
	fmt.Println()
}
