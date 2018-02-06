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
	fmt.Println(stats.StrAbility, ": ", char.AbilityScores[stats.StrAbility])
	fmt.Println(stats.DexAbility, ": ", char.AbilityScores[stats.DexAbility])
	fmt.Println(stats.ConAbility, ": ", char.AbilityScores[stats.ConAbility])
	fmt.Println(stats.IntAbility, ": ", char.AbilityScores[stats.IntAbility])
	fmt.Println(stats.WisAbility, ": ", char.AbilityScores[stats.WisAbility])
	fmt.Println(stats.ChaAbility, ": ", char.AbilityScores[stats.ChaAbility])
}

func GenerateRandomCharacter() Character {
	char := Character{}
	char.Name = "Ser Roderick"
	char.Class = classes.GetRandomClass()
	char.AbilityScores = stats.GenerateRankedScores(classes.GetRankedStatsForClass(char.Class))

	return char
}
