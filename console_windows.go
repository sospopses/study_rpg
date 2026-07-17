//go:build windows

package main

import "syscall"

func fixConsoleEncoding() {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setConsoleOutputCP := kernel32.NewProc("SetConsoleOutputCP")
	setConsoleOutputCP.Call(uintptr(65001)) // 65001 = UTF-8
}