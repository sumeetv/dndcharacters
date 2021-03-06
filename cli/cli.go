package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/sumeetv/dndcharacters/characters"
	"github.com/sumeetv/dndcharacters/classes"
	"github.com/sumeetv/dndcharacters/stats"
)

func main() {
	fmt.Println("Welcome to the D&D 5e character manager! What would you like to do?")
	fmt.Println("a - Random Attributes")
	fmt.Println("c - Random Class")
	fmt.Println("g - Random Character")
	fmt.Println("x - Exit")
	for {
		fmt.Print("Enter an option: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "a" {
			GetRandomAttributes()
		}

		if text == "c" {
			fmt.Println(classes.GetRandomClass())
		}

		if text == "g" {
			char := characters.GenerateRandomCharacter()
			char.PrintCharacter()
		}

		if text == "x" {
			break
		}
	}
}

func GetRandomAttributes() {
	total := 0
	scores := stats.GenerateAllScores()
	for key, value := range scores {
		total += value
		fmt.Println(key, ": ", value)
	}
	fmt.Println("Total: ", total)
}
