package lib

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func LoadAllHeroes(path string) []Hero {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read file: %v", err)
	}

	var heros []Hero
	err = json.Unmarshal(byteValue, &heros)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON: %v", err)
	}

	return heros
}

func LoadAllCards(path string) []Card {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read file: %v", err)
	}

	var cards []Card
	err = json.Unmarshal(byteValue, &cards)
	if err != nil {
		fmt.Printf("Failed to unmarshal JSON: %v", err)
	}

	return cards
}

type Hero struct {
	Id           int       `json:"id"`
	Class        string    `json:"class"`
	Name         string    `json:"name"`
	HeroPower    HeroPower `json:"hero_power"`
	SummonSound  string    `json:"summon_sound"`
	AttackSound  string    `json:"attack_sound"`
	DefeatSound  string    `json:"defeat_sound"`
	Health       int
	ManaCrystals int
}

type Card struct {
	Id            int           `json:"id"`
	Name          string        `json:"name"`
	ManaCost      int           `json:"mana_cost"`
	Rarity        string        `json:"rarity"`
	Description   string        `json:"description"`
	SummonSound   string        `json:"summon_sound"`
	AttackSound   string        `json:"attack_sound"`
	DeathSound    string        `json:"death_sound"`
	CardType      string        `json:"card_type"`      // "minion" || "spell" || "weapon"
	Class         string        `json:"class"`          // "mage", "warrior" ...
	MinionType    string        `json:"minion_type"`    // "beast", "neutral" ...
	MinionAbility MinionAbility `json:"minion_ability"` // Only applicable for minions
	Health        int           `json:"health"`         // Only applicable for minions
	Attack        int           `json:"attack"`         // Only applicable for minions weapons
	Durability    int           `json:"durability"`     // Only applicable for weapons
	SpellType     string        `json:"spell_type"`     // Only applicable for spells
}

type MinionAbility struct {
	Id                 int    `json:"id"`
	AbilityDescription string `json:"ability_description"`
}

type HeroPower struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Deck struct {
	Hero  Hero
	Cards []Card
}

// +------------------------------+
// | MANA: 4                      |
// |                              |
// |     Chillwind Yeti           |
// |                              |
// |   Taunt                      |
// |   4/5                        |
// |                              |
// |                              |
// |   ATTACK: 4       HEALTH: 5  |
// +------------------------------+

// +------------------------------+
// | MANA: 3                      |
// |                              |
// |     Fiery War Axe            |
// |                              |
// |   Battlecry: Equip a 3/2 Axe |
// |                              |
// |                              |
// |   ATTACK: 3       HEALTH: 2  |
// +------------------------------+

// +------------------------------+
// | MANA: 4                      |
// |                              |
// |        Fireball              |
// |                              |
// |   Deal 6 damage              |
// |   to any target.             |
// |                              |
// |   ATTACK: 0       HEALTH: 0  |
// +------------------------------+
