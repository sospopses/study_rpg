package main

import (
	"bufio"
	"fmt"
	"strings"
)

func CreateParty(reader *bufio.Reader) []Player {
	fmt.Print("Введите количество игроков: ")
	quantityPlayer := readInt(reader)
	if quantityPlayer < 1 {
		quantityPlayer = 1
	}
	var party []Player
	for i := 1; i <= quantityPlayer; i++ {
		fmt.Printf("\n=== Создание игрока №%d ===\n", i)
		hero := CreatePlayer(reader, party)
		party = append(party, hero)
	}
	return party
}
func CreatePlayer(reader *bufio.Reader, chek []Player) Player {
	var hero Player
	fmt.Println("Добро пожаловать в текстовую RPG!")

	var name string
	for {
		fmt.Print("Введите имя героя: ")
		name = readLine(reader)

		if nameError(name, chek) {
			continue
		}
		break
	}

	fmt.Println("\nВыберите ваш класс:")
	fmt.Println("1. Воин (Warrior)")
	fmt.Println("2. Маг (Mage)")
	fmt.Println("3. Паладин (Paladin)")
	fmt.Println("4. Плут (Rogue)")
	fmt.Println("5. Охотник (Hunter)")
	fmt.Print("Твой выбор: ")
	charType := readInt(reader)

	if charType < 1 || charType > 5 {
		fmt.Println("Неверный выбор! Выбран воин по умолчанию.")
		charType = 1
	}

	switch charType {
	case 1:
		hero = &Warrior{Name: name, Health: 100, MaxHP: 100, MinDamage: 12, MaxDamage: 18, Lvl: 1, XP: 0, Gold: 0}
	case 2:
		hero = &Mage{Name: name, Health: 60, MaxHP: 60, Mana: 50, MaxMana: 50, MinDamage: 8, MaxDamage: 14, Lvl: 1, XP: 0, Gold: 0}
	case 3:
		hero = &Paladin{Name: name, Health: 120, MaxHP: 120, Shield: 0, MaxShield: 50, MinDamage: 10, MaxDamage: 15, Lvl: 1, XP: 0, Gold: 0}
	case 4:
		hero = &Rogue{Name: name, Health: 75, MaxHP: 75, Energy: 100, MaxEnergy: 100, MinDamage: 14, MaxDamage: 22, Lvl: 1, XP: 0, Gold: 0}
	case 5:
		hero = &Hunter{Name: name, Health: 80, MaxHP: 80, PetHp: 40, MaxPetHp: 40, MinDamage: 10, MaxDamage: 16, Lvl: 1, XP: 0, Gold: 0}
	}

	return hero
}

func nameError(name string, party []Player) bool {
	for _, p := range party {
		if name == "" {
			fmt.Println("Имя не может быть пустым")
			return true
		}
		if strings.EqualFold(p.GetName(), name) {
			fmt.Println("Имя занято, введите дргуое")
			return true
		}

	}
	return false
}

