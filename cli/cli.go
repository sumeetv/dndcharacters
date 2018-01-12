package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the D&D 5e character manager! What would you like to do?")
	fmt.Println("g - Generate")
	fmt.Println("x - Exit")
	for {
		fmt.Print("Enter an option: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "x" {
			break
		}

		if text == "g" {
			fmt.Println("I'll find something to put here")
		}
	}
}
