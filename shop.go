package main

import (
	"bufio"
	"fmt"
)

func VisitShop(hero Player, reader *bufio.Reader) {
	healPrice := 30
	damageBaf := 10
	damageBafPrice := 75
	for {
		fmt.Println("\n--- 🏪 Магазин ---")
		fmt.Printf("Золото: %d💰\n", hero.GetGold())
		fmt.Printf("1. Полное исцеление (цена: %d)\n", healPrice)
		fmt.Printf("2. Улучшить оружие +%d урона (цена: %d)\n", damageBaf, damageBafPrice)
		fmt.Println("3. Уйти")
		action := readInt(reader)
		switch action {
		case 1:
			if hero.SpendGold(healPrice) {
				hero.FullHeal()
				fmt.Println("Ресурсы успешно восстановлены!")
			} else {
				fmt.Println("Недостаточно золота!")
			}
			pause(500)
		case 2:
			if hero.SpendGold(damageBafPrice) {
				hero.UpgradeWeapon(damageBaf)
				fmt.Println("Оружие успешно улучшено!")
			} else {
				fmt.Println("Недостаточно золота!")
			}
			pause(500)
		case 3:
			fmt.Println("Возвращайтесь ещё!")
			return
		default:
			fmt.Println("Такого действия нет.")
		}
	}
}
