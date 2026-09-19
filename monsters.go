package main

import (
	"math/rand"
)

var monsters = []Monster{
	{
		Kind:     "Чумная крыса",
		Hp:       40,
		MinPower: 4,
		MaxPower: 8,
		DropXP:   15,
		GoldDrop: 5,
		MinLvl:   1,
		Weight:   1,
	},
	{
		Kind:     "Гоблин-разбойник",
		Hp:       70,
		MinPower: 8,
		MaxPower: 12,
		DropXP:   25,
		GoldDrop: 12,
		MinLvl:   1,
		Weight:   2,
	},
	{
		Kind:     "Дикий орк",
		Hp:       110,
		MinPower: 12,
		MaxPower: 18,
		DropXP:   45,
		GoldDrop: 25,
		MinLvl:   3,
		Weight:   2,
	},
	{
		Kind:     "Проклятый скелет",
		Hp:       85,
		MinPower: 9,
		MaxPower: 14,
		DropXP:   30,
		GoldDrop: 15,
		MinLvl:   2,
		Weight:   1,
	},
	{
		Kind:     "Пещерный тролль (МИНИ-БОСС)",
		Hp:       180,
		MinPower: 16,
		MaxPower: 24,
		DropXP:   100,
		GoldDrop: 60,
		MinLvl:   5,
		Weight:   10,
	},
}

// GetRandomMonster генерирует случайного противника на основе структуры из types.go
func GetRandomMonster(playerLvl int) *Monster {
	var availableMonsters []Monster
	for _, m := range monsters {
		if playerLvl >= m.MinLvl {
			availableMonsters = append(availableMonsters, m)
		}
	}

	if len(availableMonsters) == 0 {
		panic("нет доступных монстров для уровня героя")
	}

	m := PickWeightedMonster(availableMonsters)
	return m
}

// PickWeightedMonster выбирает монстра случайно, где МЕНЬШИЙ Weight = ВЫШЕ шанс выпадения.
func PickWeightedMonster(candidates []Monster) *Monster {
	effWeights := make([]float64, len(candidates))
	var sumWeight float64

	for i, m := range candidates {
		effWeights[i] = 1.0 / float64(m.Weight)
		sumWeight += effWeights[i]
	}

	r := rand.Float64() * sumWeight
	for i, w := range effWeights {
		r -= w
		if r < 0 {
			return &candidates[i]
		}
	}

	// Подстраховка от погрешности float64 — без явного return в конце
	// Go не скомпилирует функцию (не все пути возвращают значение).
	return &candidates[len(candidates)-1]
}

func GenerateMonsterParty(party []Player) []Monster {
	var monster []Monster
	for i := 0; i < len(party); i++ {
		monster = append(monster, *GetRandomMonster(PartyLvl(party)))
	}
	return monster
}

func IsMonstersAlive(monsters []Monster) bool {
	for _, m := range monsters {
		if m.Hp > 0 {
			return true
		}
	}
	return false
}