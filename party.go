package main

import (
//	"bufio"
	"fmt"
)

func Party_Stats(party []Player) {
	fmt.Printf("\nСостояние команды:\n")
	for _, p := range party {
		p.ShowStats()
		if p.GetHealth() > 0 {
			fmt.Printf("| Состояние: Жив\n")
		} else {
			fmt.Printf("| Состояние: Мёртв\n")
		}
	}
	fmt.Println("")
}

func IsPartyAlive (party []Player) bool {
	for _, p := range party {
		if p.GetHealth() > 0 {
			return true
		}
	}
	return false
}

func PartyLvl (party []Player) int {
	sum_lvl := 0
	for _, p := range party {
		sum_lvl += p.GetLevel()
	}
	return int(float64(sum_lvl)/float64(len(party)))
}