package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// runOne()
	// runTwo()
	// runThree()
	runFour()
}

func runOne() {
	c := make(chan string)
	go boring1("Boring1.. ", c)
	fmt.Println("im listening")
	for i := 0; i < 5; i++ {
		fmt.Printf("you say %q\n", <-c)
	}
	fmt.Println("you are boring, imleaving")
}

func boring1(text string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("%s %d", text, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func runTwo() {
	c := boring2("Boring2")
	fmt.Println("Imlistening")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say %s\n", <-c)
	}
	fmt.Println("you're boring, im leaving")
}

func boring2(text string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- fmt.Sprintf("%s %d", text, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func runThree() {
	joe := boring2("Joe")
	ann := boring2("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("you are both boring, im leaving")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func runFour() {
	c := fanIn(boring2("Joe"), boring2("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("you are both boring, im leaving")
}
