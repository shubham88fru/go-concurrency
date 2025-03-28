package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

var hunger = 3 //how many times each philosopher will eat
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

var orderMutex sync.Mutex
var orderFinished []string

func main() {
	// print out a welcome message
	fmt.Println("Dining Philosopher's Problem")
	fmt.Println("===================================")
	fmt.Println("The table is empty.")

	time.Sleep(sleepTime)

	//start a meal
	dine()

	//print out finished message
	fmt.Println("The Table is empty.")

}

func dine() {
	done := &sync.WaitGroup{}
	done.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		go dinningProblem(philosophers[i], forks, done, seated)
	}

	done.Wait()
}

func dinningProblem(philosopher Philosopher, forks map[int]*sync.Mutex, done *sync.WaitGroup, seated *sync.WaitGroup) {
	defer done.Done()

	fmt.Printf("%s is seated at the table.\n", philosopher.name)
	seated.Done()

	seated.Wait()

	for i := hunger; i > 0; i-- {

		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s picked up the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s picked up the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("\t%s picked up the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("\t%s picked up the right fork.\n", philosopher.name)
		}

		fmt.Printf("\t%s is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.rightFork].Unlock()
		forks[philosopher.leftFork].Unlock()
		fmt.Printf("\t%s put down the the forks.\n", philosopher.name)
	}

	fmt.Printf("%s is done eating and is leaving the table.\n", philosopher.name)

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.name)
	orderMutex.Unlock()
}
