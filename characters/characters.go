package characters

import (
	"fmt"

	"github.com/sumeetv/dndcharacters/classes"
	"github.com/sumeetv/dndcharacters/stats"
)

type Character struct {
	Name          string
	AbilityScores map[stats.AbilityScore]int
  SavingThrows  map[stats.AbilityScore]int
	Class         classes.PlayableClass
}

func (char Character) PrintCharacter() {
	fmt.Println("Name: ", char.Name)
	fmt.Println("Class: ", char.Class)

  abilities := stats.GetAbilityScoreList()

	fmt.Println("Ability Scores:")
  for _, ability := range abilities {
    fmt.Println(ability, ": ", char.AbilityScores[ability])
  }

  fmt.Println("Saving Throws")
  for _, ability := range abilities {
    fmt.Println(ability, ": ", char.SavingThrows[ability])
  }
}

func GenerateRandomCharacter() Character {
	char := Character{}
	char.Name = "Ser Roderick"
	char.Class = classes.GetRandomClass()
	char.AbilityScores = stats.GenerateRankedScores(classes.GetRankedStatsForClass(char.Class))

  // Refresh the character to generate calculated stats
  char.RefreshCharacter()
	return char
}

// Recalculate all calculated values from scratch
func (char *Character)RefreshCharacter() {
  abilities := stats.GetAbilityScoreList()

  // Saving Throws
  if (char.SavingThrows == nil) {
    char.SavingThrows = make(map[stats.AbilityScore]int)
  }
  for _, ability := range abilities {
    char.SavingThrows[ability] = stats.CalculateSavingThrow(char.AbilityScores[ability])
  }
}
