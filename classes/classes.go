package classes

import (
	"math/rand"
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
