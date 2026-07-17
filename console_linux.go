//go:build !windows

package main

func fixConsoleEncoding() {
    // В Linux терминалы по умолчанию работают в UTF-8, 
    // поэтому здесь ничего делать не нужно.
}