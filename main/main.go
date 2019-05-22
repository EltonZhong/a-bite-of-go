package main

import "fmt"

func main() {
	hello("world")
}

type Monkey interface {
	eat() string
}

type Bird interface {
	eat() string
}

type Monkee struct {
}

func init() {
	fmt.Println("This will be inited")
}

func hello(str string) {
	var a = "hello " + str
	fmt.Print(a)

	var b = Monkee{}
	fmt.Println(banana(b))
	fmt.Println(worm(b))
}

func (Monkee) eat() string {
	return "hello"
}

func banana(animal Monkey) string {
	return animal.eat()
}

func worm(animal Bird) string {
	return animal.eat()
}
