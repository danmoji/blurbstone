package tests

import (
	"blurbstone/lib"
	"reflect"
	"testing"
)

func TestLoadHeros(t *testing.T) {
	expected := []lib.Hero{
		{
			Id:    1,
			Class: "mage",
			Name:  "Jaina Proudmoreee",
			HeroPower: lib.HeroPower{
				Id:          1,
				Name:        "Fireball",
				Description: "Deals 1 damage.",
			},
			SummonSound: "You've asked for it",
			AttackSound: "You've asked for it",
			DefeatSound: "AAAARRRH!!!!!!",
		},
		{
			Id:    2,
			Class: "paladin",
			Name:  "Tirion Fordring",
			HeroPower: lib.HeroPower{
				Id:          2,
				Name:        "Reinforce",
				Description: "Summon a 1/1 Silver Hand Recruit.",
			},
			SummonSound: "For Honor",
			AttackSound: "Justice aaa...",
			DefeatSound: "Justice aaa...",
		},
	}

	// Call LoadHeros with the path to `./heros.json`
	result := lib.LoadAllHeroes("./heroes-mock.json")

	// Compare the result with the expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("LoadHeros() = %v, want %v", result, expected)
	}
}

func TestLoadCards(t *testing.T) {
	// Define the expected result with data that matches the provided cards-test.json content.
	expect := []lib.Card{
		{
			Id:          1,
			Name:        "Chillwind Yeti",
			ManaCost:    4,
			Rarity:      "free",
			Description: "",
			SummonSound: "AAAARRRH",
			AttackSound: "AAAARRRH",
			DeathSound:  "AAAARRRH",
			CardType:    "minion",
			Class:       "neutral",
			MinionType:  "neutral",
			MinionAbility: lib.MinionAbility{
				Id:                 0,
				AbilityDescription: "",
			},
			Health:     5,
			Attack:     4,
			Durability: 0,
			SpellType:  "",
		},
		{
			Id:          2,
			Name:        "Fire ball",
			ManaCost:    4,
			Rarity:      "free",
			Description: "Deals 6 damage.",
			SummonSound: "AAAARRRH",
			AttackSound: "",
			DeathSound:  "",
			CardType:    "spell",
			Class:       "mage",
			MinionType:  "",
			MinionAbility: lib.MinionAbility{
				Id:                 0,
				AbilityDescription: "",
			},
			Health:     0,
			Attack:     0,
			Durability: 0,
			SpellType:  "damage",
		},
		{
			Id:          3,
			Name:        "Fiery waraxe",
			ManaCost:    2,
			Rarity:      "free",
			Description: "",
			SummonSound: "ZINK!",
			AttackSound: "ZINK!",
			DeathSound:  "ZINK!",
			CardType:    "weapon",
			Class:       "warrior",
			MinionType:  "",
			MinionAbility: lib.MinionAbility{
				Id:                 0,
				AbilityDescription: "",
			},
			Health:     0,
			Attack:     0,
			Durability: 0,
			SpellType:  "",
		},
		{
			Id:          4,
			Name:        "Wisp",
			ManaCost:    0,
			Rarity:      "common",
			Description: "",
			SummonSound: "ZINK!",
			AttackSound: "ZINK!",
			DeathSound:  "ZINK!",
			CardType:    "minion",
			Class:       "neutral",
			MinionType:  "undead",
			MinionAbility: lib.MinionAbility{
				Id:                 0,
				AbilityDescription: "",
			},
			Health:     1,
			Attack:     1,
			Durability: 0,
			SpellType:  "",
		},
		{
			Id:          5,
			Name:        "Shieldbearer",
			ManaCost:    1,
			Rarity:      "common",
			Description: "Taunt",
			SummonSound: "You shall not pass!",
			AttackSound: "You shall not pass!",
			DeathSound:  "You shall not pass!",
			CardType:    "minion",
			Class:       "neutral",
			MinionType:  "neutral",
			MinionAbility: lib.MinionAbility{
				Id:                 1,
				AbilityDescription: "Taunt",
			},
			Health:     4,
			Attack:     0,
			Durability: 0,
			SpellType:  "",
		},
		{
			Id:          6,
			Name:        "Worgen Infiltrator",
			ManaCost:    1,
			Rarity:      "common",
			Description: "Stealth",
			SummonSound: "I smell blood!",
			AttackSound: "I smell blood!",
			DeathSound:  "I smell blood!",
			CardType:    "minion",
			Class:       "neutral",
			MinionType:  "neutral",
			MinionAbility: lib.MinionAbility{
				Id:                 2,
				AbilityDescription: "Stealth",
			},
			Health:     1,
			Attack:     2,
			Durability: 0,
			SpellType:  "",
		},
		{
			Id:          7,
			Name:        "Stonetusk Boar",
			ManaCost:    1,
			Rarity:      "common",
			Description: "Charge",
			SummonSound: "Kroch Kroch Kroch!",
			AttackSound: "Kroch Kroch Kroch!",
			DeathSound:  "Kroch Kroch Kroch!",
			CardType:    "minion",
			Class:       "neutral",
			MinionType:  "neutral",
			MinionAbility: lib.MinionAbility{
				Id:                 3,
				AbilityDescription: "Charge",
			},
			Health:     1,
			Attack:     1,
			Durability: 0,
			SpellType:  "",
		},
	}

	// Load cards from the JSON test file
	result := lib.LoadAllCards("cards-mock.json")

	// Check if the loaded cards match the expected output
	if !reflect.DeepEqual(result, expect) {
		t.Errorf("Expected %v, but got %v", expect, result)
	}
}
