package abilityscore

import (
	"math/rand"
	"sort"
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

func GenerateAllScores() map[string]int {
	scores := map[string]int{
		"STR": GenerateAbilityScore(),
		"DEX": GenerateAbilityScore(),
		"CON": GenerateAbilityScore(),
		"INT": GenerateAbilityScore(),
		"WIS": GenerateAbilityScore(),
		"CHA": GenerateAbilityScore(),
	}
	return scores
}
