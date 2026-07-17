package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fixConsoleEncoding()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("🧙 Добро пожаловать в текстовую RPG!")
	fmt.Print("Введите имя героя: ")
	name := readLine(reader)

	fmt.Println("\nВыберите ваш класс:")
	fmt.Println("1. 🛡️ Воин (Warrior)")
	fmt.Println("2. 🔮 Маг (Mage)")
	fmt.Println("3. 👑 Паладин (Paladin)")
	fmt.Println("4. 🗡️ Плут (Rogue)")
	fmt.Println("5. 🏹 Охотник (Hunter)")
	fmt.Print("Твой выбор: ")
	charType := readInt(reader)

	var hero Player

	switch charType {
	case 1:
		hero = &Warrior{Name: name, Health: 100, MaxHP: 100, MinDamage: 12, MaxDamage: 18, Lvl: 1, XP: 0, Gold: 0}
	case 2:
		hero = &Mage{Name: name, Health: 60, MaxHP: 60, Mana: 50, MaxMana: 50, MinDamage: 8, MaxDamage: 14, Lvl: 1, XP: 0, Gold: 0}
	case 3:
		hero = &Paladin{Name: name, Health: 120, MaxHP: 120, Shield: 0, MaxShield: 50, MinDamage: 10, MaxDamage: 15, Lvl: 1, XP: 0, Gold: 0}
	case 4:
		hero = &Rogue{Name: name, Health: 75, MaxHP: 75, Energy: 100, MinDamage: 14, MaxDamage: 22, Lvl: 1, XP: 0, Gold: 0}
	case 5:
		hero = &Hunter{Name: name, Health: 80, MaxHP: 80, PetHp: 50, MinDamage: 10, MaxDamage: 16, Lvl: 1, XP: 0, Gold: 0}
	default:
		fmt.Println("Неверный выбор! Выбран воин по умолчанию.")
		hero = &Warrior{Name: name, Health: 100, MaxHP: 100, MinDamage: 12, MaxDamage: 18, Lvl: 1, XP: 0, Gold: 0}
	}

	// --- ИГРОВОЙ ЦИКЛ ---
	for {
		// Проверка смерти
		if hero.GetHealth() <= 0 {
			fmt.Println("\n💀 Ты погиб... Твое приключение окончено.\nНажмите Enter чтобы выйти")
			readLine(reader) // просто ждём Enter, значение не нужно
			break
		}

		fmt.Println("\n--- Меню ---")
		fmt.Println("1. Искать монстра")
		fmt.Println("2. Посмотреть статы")
		fmt.Println("3. Магазин")
		fmt.Println("4. Выход")
		fmt.Print("Выбери действие: ")

		action := readInt(reader)

		switch action {
		case 1:
			// Генерируем монстра и запускаем бой
			monster := GetRandomMonster()
			StartCombat(hero, monster, reader)
		case 2:
			hero.ShowStats()
		case 3:
			VisitShop(hero, reader)
		case 4:
			fmt.Println("До встречи!")
			return
		default:
			fmt.Println("Такого действия нет.")
		}
	}
}
