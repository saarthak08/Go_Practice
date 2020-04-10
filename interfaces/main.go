package main

import (
	"fmt"
)

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting(string) string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(sb, "hello")
	printGreeting(eb, "hello")
}

func (e englishBot) getGreeting(abc string) string {
	// A very custom different logic
	return "English Bot Here!"
}

func (s spanishBot) getGreeting(abc string) string {
	return "Spanish Bot Here!"
}

func printGreeting(b bot, abc string) {
	fmt.Println(b.getGreeting(abc))
}
