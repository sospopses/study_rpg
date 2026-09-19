package main

import (
	"bufio"
	"fmt"
)

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

func SelectTarget (monsters *[]Monster, reader *bufio.Reader) *Monster {
	if len(*monsters) == 1 {
		return &(*monsters)[0]
	}

	fmt.Printf("\nЖивые монстры:\n")

	for i, m := range *monsters {
		if m.Hp > 0 {
			fmt.Printf("%d. %s | ХП: %d | Урон: %d-%d\n", (i+1), m.Kind, m.Hp, m.MinPower, m.MaxPower)
		}
	}

	fmt.Println("Выберите монстра(введите номер):")

	for {
		target := readInt(reader)
		if target < 1 || target > len(*monsters) {
			fmt.Printf("Неверный выбор, выберите от 1 до %d\n", len(*monsters))
			continue
		}
		if (*monsters)[target-1].Hp <=0 {
			fmt.Println("Неверный выбор, этот монстр мёртв")
			continue
		}
		return &(*monsters)[target-1]
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
			monster := GenerateMonsterParty(party)
			StartCombat(party, &monster, reader)
			pause_midle()
		case 2:
			Party_Stats(party)
		case 3:
    		fmt.Println("\nКого отправим за покупками?")
    		buyer := SelectPlayer(party, reader)
    		VisitShop(buyer, reader)
		case 4:
			fmt.Println("До встречи!")
			return
		default:
			fmt.Println("Такого действия нет.")
		}
	}
}
