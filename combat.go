package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
)

func StartCombat(player Player, monster *Monster, reader *bufio.Reader) {
	fmt.Printf("\nНАЧАЛСЯ БОЙ! %s сталкивается лицом к лицу с %s!\n", player.GetName(), monster.Kind)
	pause_midle()

	for player.GetHealth() > 0 && monster.Hp > 0 {
		player.ShowStats()
		fmt.Printf("%s | HP: %d | Урон: %d-%d\n", monster.Kind, monster.Hp, monster.MinPower, monster.MaxPower)
		pause_short()

		// --- ХОД ИГРОКА ---
		playerDmg := player.GetDamage(reader)
		monster.Hp -= playerDmg
		fmt.Printf("Вы нанесли %s %d урона! (Осталось HP монстра: %d)\n", monster.Kind, playerDmg, monster.Hp)
		pause_midle()

		if monster.Hp <= 0 {
			fmt.Printf("\nПобеда! %s повержен!\n", monster.Kind)
			pause_midle()
			player.AddGold(monster.GoldDrop)

			requiredXP := int(100 * math.Pow(1.5, float64(player.GetLevel()-1)))
			leveledUp := player.AddXP(monster.DropXP, requiredXP)
			fmt.Printf("Получено золота: %d | Получено опыта: %d\n", monster.GoldDrop, monster.DropXP)
			pause_midle()

			if leveledUp {
				newHP := player.LevelUp()
				fmt.Printf("УРОВЕНЬ ПОВЫШЕН! Твое максимальное здоровье теперь: %d HP!\n", newHP)
				pause_midle()
			}
			return
		}

		// --- ХОД МОНСТРА ---
		fmt.Printf("\nХод врага: %s замахивается...\n", monster.Kind)
		pause_midle()

		// Считаем случайный урон монстра
		monsterDmg := monster.MinPower
		if monster.MaxPower > monster.MinPower {
			monsterDmg = rand.Intn(monster.MaxPower-monster.MinPower+1) + monster.MinPower
		}

		currentHP := player.GetHealth()
		player.SetHealth(currentHP - monsterDmg)
		pause_midle()

		if player.GetHealth() <= 0 {
			fmt.Printf("\nВы погибли в бою с %s... Игра окончена.\n", monster.Kind)
			return
		}
		fmt.Println("-------------------------------------------")
		pause_short()
	}
}
