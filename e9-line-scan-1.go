package main

import (
	"bufio"
	"fmt"
	"os"
)

//once
func main0() {

	fmt.Print("Enter text: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	//fmt.Println(input)
	fmt.Println("Input string: ", input.Text())
}

//multiple
func main2() {
	//fmt.Print("Enter text: ")
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		input.Scan()
		fmt.Println("Input string: ", input.Text())
	}
}

//multiple + condition
func main() {
	fmt.Print("Enter text: ")
	c := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println("Input string: ", input.Text())
		fmt.Print("Enter text: ")
		c[input.Text()]++

		if input.Text() == "break" {
			break
		}
	}

	fmt.Println(c)
}

//https://go.dev/play/p/W0ZiHOT6kfr