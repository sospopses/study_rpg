package main

import (
	"bufio"
	"fmt"
)

func IsPartyAlive (party []Player) bool {
	for _, p := range party {
		if p.GetHealth() > 0 {
			return true
		}
	}
	return false
}

func SelectPlayer (party []Player, reader *bufio.Reader) Player {
	if len(party) == 1 {
		return party[0]
	}
	
	for i, p := range party {
		fmt.Printf("%d. %s, ХП: %d\n", (i+1), p.GetName(), p.GetHealth())
	}

	fmt.Println("Введите номер героя:")

	for {
		hero := readInt(reader)
		if hero < 1 || hero > len(party)  {
			fmt.Printf("Неверный выбор, выберите номер от 1 до %d\n", len(party))
			continue
		}
		if party[hero-1].GetHealth() <= 0 {	
			fmt.Println("Герой мёрв, выберите другого")
			continue
		}
		return party[hero-1]
	}
}

func GameCycle(party []Player, reader *bufio.Reader) {
	for {
		// Проверка смерти
		if !IsPartyAlive(party) {
			fmt.Println("\nВаша команда проиграла.\nНажмите Enter чтобы выйти")
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
			monster := GenerateMonsterParty(party)[0]
			StartCombat(party[0], &monster, reader)
		case 2:
			fmt.Println("\n=== СТАТУС ОТРЯДА ===")
			for _, h := range party {
				h.ShowStats()
			}
		case 3:
			VisitShop(party[0], reader)
		case 4:
			fmt.Println("До встречи!")
			return
		default:
			fmt.Println("Такого действия нет.")
		}
	}
}
