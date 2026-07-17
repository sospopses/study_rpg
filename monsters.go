package main

import (
	"math/rand"
)

// GetRandomMonster генерирует случайного противника на основе структуры из types.go
func GetRandomMonster() *Monster {
	monsters := []Monster{
		{
			Kind:     "Чумная крыса",
			Hp:       40,
			MinPower: 4,
			MaxPower: 8,
			DropXP:   15,
			GoldDrop: 5,
		},
		{
			Kind:     "Гоблин-разбойник",
			Hp:       70,
			MinPower: 8,
			MaxPower: 12,
			DropXP:   25,
			GoldDrop: 12,
		},
		{
			Kind:     "Дикий орк",
			Hp:       110,
			MinPower: 12,
			MaxPower: 18,
			DropXP:   45,
			GoldDrop: 25,
		},
		{
			Kind:     "Проклятый скелет",
			Hp:       85,
			MinPower: 9,
			MaxPower: 14,
			DropXP:   30,
			GoldDrop: 15,
		},
		{
			Kind:     "Пещерный тролль (МИНИ-БОСС)",
			Hp:       180,
			MinPower: 16,
			MaxPower: 24,
			DropXP:   100,
			GoldDrop: 60,
		},
	}

	randomIndex := rand.Intn(len(monsters))
	m := monsters[randomIndex]
	return &m
}
