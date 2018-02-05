package characters

import (
	"fmt"

	"github.com/sumeetv/dndcharacters/classes"
	"github.com/sumeetv/dndcharacters/stats"
)

type Character struct {
	Name          string
	AbilityScores map[stats.AbilityScore]int
	Class         string
}

func (char Character) PrintCharacter() {
	fmt.Println("Name: ", char.Name)
	fmt.Println("Class: ", char.Class)

	fmt.Println("Ability Scores:")
	for score, value := range char.AbilityScores {
		fmt.Println(score, ": ", value)
	}
}

func GenerateRandomCharacter() Character {
	char := Character{}
	char.Name = "Ser Roderick"
	char.Class = classes.GetRandomClass()
	char.AbilityScores = stats.GenerateRankedScores(classes.GetRankedStatsForClass(char.Class))

	return char
}
