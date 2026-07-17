package main

import (
	"bufio"
	"fmt"
	"math/rand"
)

func StartCombat(player Player, monster *Monster, reader *bufio.Reader) {
	fmt.Printf("\n⚔️ НАЧАЛСЯ БОЙ! %s сталкивается лицом к лицу с %s! ⚔️\n", player.GetName(), monster.Kind)
	pause(700)

	for player.GetHealth() > 0 && monster.Hp > 0 {
		player.ShowStats()
		fmt.Printf("😈 %s | HP: %d | Урон: %d-%d\n", monster.Kind, monster.Hp, monster.MinPower, monster.MaxPower)
		pause(400)

		// --- ХОД ИГРОКА ---
		playerDmg := player.GetDamage(reader)
		monster.Hp -= playerDmg
		fmt.Printf("💥 Вы нанесли %s %d урона! (Осталось HP монстра: %d)\n", monster.Kind, playerDmg, monster.Hp)
		pause(700)

		if monster.Hp <= 0 {
			fmt.Printf("\n🎉 Победа! %s повержен!\n", monster.Kind)
			pause(600)
			player.AddGold(monster.GoldDrop)

			leveledUp := player.AddXP(monster.DropXP, 100)
			fmt.Printf("💰 Получено золота: %d | ✨ Получено опыта: %d\n", monster.GoldDrop, monster.DropXP)
			pause(600)

			if leveledUp {
				newHP := player.LevelUp()
				fmt.Printf("🆙 УРОВЕНЬ ПОВЫШЕН! Твое максимальное здоровье теперь: %d HP!\n", newHP)
				pause(700)
			}
			return
		}

		// --- ХОД МОНСТРА ---
		fmt.Printf("\n🥊 Ход врага: %s замахивается...\n", monster.Kind)
		pause(600)

		// Считаем случайный урон монстра
		monsterDmg := monster.MinPower
		if monster.MaxPower > monster.MinPower {
			monsterDmg = rand.Intn(monster.MaxPower-monster.MinPower+1) + monster.MinPower
		}

		currentHP := player.GetHealth()
		player.SetHealth(currentHP - monsterDmg)
		pause(600)

		if player.GetHealth() <= 0 {
			fmt.Printf("\n💀 Вы погибли в бою с %s... Игра окончена.\n", monster.Kind)
			return
		}
		fmt.Println("-------------------------------------------")
		pause(300)
	}
}
