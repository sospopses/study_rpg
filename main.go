package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	party := CreateParty(reader)
	GameCycle(party, reader)
}
