package stats

import (
	"math/rand"
	"sort"
)

type AbilityScore string

const (
	StrAbility = "STR"
	DexAbility = "DEX"
	ConAbility = "CON"
	IntAbility = "INT"
	WisAbility = "WIS"
	ChaAbility = "CHA"
)

func GenerateAbilityScore() int {
	diceRolls := []int{
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
		rand.Intn(6) + 1,
	}
	sort.Ints(diceRolls)

	// Return the total of the three highest rolls
	return diceRolls[1] + diceRolls[2] + diceRolls[3]
}

func GenerateAllScores() map[AbilityScore]int {
	scores := map[AbilityScore]int{
		StrAbility: GenerateAbilityScore(),
		DexAbility: GenerateAbilityScore(),
		ConAbility: GenerateAbilityScore(),
		IntAbility: GenerateAbilityScore(),
		WisAbility: GenerateAbilityScore(),
		ChaAbility: GenerateAbilityScore(),
	}
	return scores
}
