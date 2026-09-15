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
	if line == "STOP" {
		panic("СТОП!!!")
	}
	return strings.TrimSpace(line)
}

// readInt читает строку и пытается превратить её в число.
// Если пользователь ввёл не число — просит повторить ввод, а не падает
func readInt(reader *bufio.Reader) int {
	for {
		line := readLine(reader)
		if line == "STOP" {
			panic("СТОП!!!")
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Print("Введите число: ")
			continue
		}
		return num
	}
}

// pause делает паузу на заданное число миллисекунд — для драматичности вывода
//func pause(ms int) {
//	time.Sleep(time.Duration(ms) * time.Millisecond)
//}

func pause_short() {
	time.Sleep(time.Duration(300) * time.Millisecond)
}

func pause_midle() {
	time.Sleep(time.Duration(600) * time.Millisecond)
}

func pause_long() {
	time.Sleep(time.Duration(900) * time.Millisecond)
}
