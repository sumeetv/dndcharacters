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
	NumScores  = 6
)

func GetAbilityScoreList() []AbilityScore {
  list := []AbilityScore{
    StrAbility,
    DexAbility,
    ConAbility,
    IntAbility,
    WisAbility,
    ChaAbility,
  }
  return list
}

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

func GenerateRankedScores(rankedScores []AbilityScore) map[AbilityScore]int {
	// TODO: Validate length of input or add support for having throwaway stats
	var scores []int
	for i := 0; i < NumScores; i++ {
		scores = append(scores, GenerateAbilityScore())
	}
	sort.Sort(sort.Reverse(sort.IntSlice(scores)))
	abilityScores := map[AbilityScore]int{}
	for index, ability := range rankedScores {
		abilityScores[ability] = scores[index]
	}
	return abilityScores
}

func CalculateSavingThrow(score int) int {
  return (score - 12) / 2
}
