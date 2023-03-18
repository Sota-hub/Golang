package main

import "fmt"


// add new type called bot to all of the different type in the program
type bot interface {
 // a type in this program that call getGreeting and return a string then it is one of the bot (= eb and sb are member of bot)
  getGreeting() string

} 

type englishBot struct {}
type spanishBot struct {}


func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string { // Associate by receiver.  Here should match to the content in bot interface
	// very custom logic for generating an english greeting
	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}