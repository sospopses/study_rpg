package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// readLine читает строку целиком (с пробелами), убирая \n/\r и пробелы по краям
func readLine(reader *bufio.Reader) string {
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

// readInt читает строку и пытается превратить её в число.
// Если пользователь ввёл не число — просит повторить ввод, а не падает
func readInt(reader *bufio.Reader) int {
	for {
		line := readLine(reader)
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Print("❌ Введите число: ")
			continue
		}
		return num
	}
}

// pause делает паузу на заданное число миллисекунд — для драматичности вывода
func pause(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}
