package main

import "fmt"

const LoginToken string = "sadfhsjfsj" // Public variable

func main() {
	var username string = "Chaitu"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)
	var smallFloat float32 = 255.5456456465456465
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)
	var largeFloat float64 = 255.5456456465456465
	fmt.Println(largeFloat)
	fmt.Printf("Variable is of type: %T \n", largeFloat)
	// default values and some aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Another variable is of type: %T \n", anotherVar)
	var stringVar string
	fmt.Println(stringVar)
	fmt.Printf("Another variable is of type: %T \n", stringVar)
	// implicit type
	var google = "http://google.com"
	fmt.Println(google)
	myWebURL := "http://coderchaitu.com"
	fmt.Println(myWebURL)

}
