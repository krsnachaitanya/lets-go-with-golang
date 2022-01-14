package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the userInput"
	fmt.Println(welcome)

	// comma ok || err ok
	fmt.Println("Enter your pizza name:")
	pizzaName, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Printf("Enter rating for %v:", pizzaName)
	rating, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Printf("Thanks for rating %v for %v", rating, pizzaName)
}
