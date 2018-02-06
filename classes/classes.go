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
	if class == "Barbarian" {
		rankedScores = []stats.AbilityScore{
			stats.ConAbility,
			stats.StrAbility,
			stats.DexAbility,
			stats.ChaAbility,
			stats.WisAbility,
			stats.IntAbility,
		}
	}
	if class == "Bard" {
		rankedScores = []stats.AbilityScore{
			stats.ChaAbility,
			stats.DexAbility,
			stats.ConAbility,
			stats.WisAbility,
			stats.StrAbility,
			stats.IntAbility,
		}
	}
	if class == "Cleric" {
		rankedScores = []stats.AbilityScore{
			stats.WisAbility,
			stats.StrAbility,
			stats.ConAbility,
			stats.ChaAbility,
			stats.DexAbility,
			stats.IntAbility,
		}
	}
	if class == "Druid" {
		rankedScores = []stats.AbilityScore{
			stats.WisAbility,
			stats.DexAbility,
			stats.ConAbility,
			stats.ChaAbility,
			stats.IntAbility,
			stats.StrAbility,
		}
	}
	if class == "Fighter" {
		rankedScores = []stats.AbilityScore{
			stats.StrAbility,
			stats.ConAbility,
			stats.DexAbility,
			stats.ChaAbility,
			stats.WisAbility,
			stats.IntAbility,
		}
	}
	if class == "Monk" {
		rankedScores = []stats.AbilityScore{
			stats.DexAbility,
			stats.WisAbility,
			stats.ConAbility,
			stats.ChaAbility,
			stats.IntAbility,
			stats.StrAbility,						
		}
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
	if class == "Ranger" {
		rankedScores = []stats.AbilityScore{
			stats.DexAbility,
			stats.WisAbility,
			stats.ConAbility,
			stats.ChaAbility,
			stats.IntAbility,
			stats.StrAbility,
		}
	}
	if class == "Rogue" {
		rankedScores = []stats.AbilityScore{
			stats.DexAbility,
			stats.ChaAbility,
			stats.StrAbility,
			stats.ConAbility,
			stats.WisAbility,
			stats.IntAbility,
		}
	}
	if class == "Sorcerer" {
		rankedScores = []stats.AbilityScore{
			stats.ChaAbility,
			stats.DexAbility,
			stats.ConAbility,
			stats.WisAbility,
			stats.IntAbility,
			stats.StrAbility,
		}
	}
	if class == "Warlock" {
		rankedScores = []stats.AbilityScore{
			stats.ChaAbility,
			stats.ConAbility,
			stats.DexAbility,
			stats.WisAbility,
			stats.StrAbility,
			stats.IntAbility,
		}
	}
	if class == "Wizard" {
		rankedScores = []stats.AbilityScore{
			stats.IntAbility,
			stats.DexAbility,
			stats.ConAbility,
			stats.WisAbility,
			stats.ChaAbility,
			stats.StrAbility,
		}
	}

	return rankedScores
}
