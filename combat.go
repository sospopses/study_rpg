package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
)

func StartCombat(party []Player, monsters *[]Monster, reader *bufio.Reader) {
	fmt.Println("\nБОЙ НАЧАЛСЯ!")
	pause_short()
	Party_Stats(party)

	for IsPartyAlive(party) && IsMonstersAlive(*monsters) {

		// --- ФАЗА ИГРОКОВ ---
		for _ , player := range party {
			if player.GetHealth() <=0 {
				continue
			} 
			fmt.Printf("\n Ходит %s", player.GetName())
			pause_short()
			monster := SelectTarget(monsters, reader)
			pause_midle()
			damage := player.GetDamage(reader)
			pause_midle()
			monster.Hp -= damage
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
			}
			if !IsMonstersAlive(*monsters) {
				fmt.Printf("Бой окончен, победа!")
				return
			}
		}

		fmt.Printf("")
		pause_midle()

		// --- ФАЗА МОНСТРОВ ---
		for _, monster := range *monsters {
			if monster.Hp <= 0 {
				continue
			}

			var alivePlayers []Player
			for _, p := range party {
				if p.GetHealth() > 0 {
					alivePlayers = append(alivePlayers, p)
				}
			}

			if len(alivePlayers) == 0 {
				break 
			}

			targetPlayer := alivePlayers[rand.Intn(len(alivePlayers))]

			monsterDmg := monster.MinPower
			if monster.MaxPower > monster.MinPower {
				monsterDmg = rand.Intn(monster.MaxPower-monster.MinPower+1) + monster.MinPower
			}
			fmt.Printf("%s атакует %s и наносит %d урона!\n", monster.Kind, targetPlayer.GetName(), monsterDmg)
			currentHP := targetPlayer.GetHealth()
			targetPlayer.SetHealth(currentHP - monsterDmg)
			pause_midle()
		}	
	}
}