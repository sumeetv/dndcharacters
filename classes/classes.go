package classes

import (
	"math/rand"

	"github.com/sumeetv/dndcharacters/stats"
)

func GetClasses() []string {
	classes := []string{
		"Barbarian",
		"Bard",
		"Cleric",
		"Druid",
		"Fighter",
		"Monk",
		"Paladin",
		"Ranger",
		"Rogue",
		"Sorcerer",
		"Warlock",
		"Wizard",
	}
	return classes
}

func GetRandomClass() string {
	classes := GetClasses()
	return classes[rand.Intn(len(classes))]
}

func GetRankedStatsForClass(class string) []stats.AbilityScore {
	rankedScores := []stats.AbilityScore{
		stats.StrAbility,
		stats.DexAbility,
		stats.ConAbility,
		stats.IntAbility,
		stats.WisAbility,
		stats.ChaAbility,
	}
	if class == "Paladin" {
		rankedScores = []stats.AbilityScore{
			stats.StrAbility,
			stats.ConAbility,
			stats.ChaAbility,
			stats.DexAbility,
			stats.WisAbility,
			stats.IntAbility,
		}
	}

	return rankedScores
}
