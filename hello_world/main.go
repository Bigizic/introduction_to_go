package main

import "fmt"

/**
 * main - simple function that prints hello world to stdout
*/
func main() {
	fmt.Println("Hello, World")

	// cals other imports here
	QuoteHelloWorldEntry();
}


/**
 * QuoteHelloWorldEntry - a function that imports the QuoteHelloWorld function
 * from the quote_hello_world package and calls it
*/

func QuoteHelloWorldEntry() {
	QuoteHelloWorld();
}
