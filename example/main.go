package main

import (
	"colorize"
	"log"
	"os"
)

func writeToStdout() {
	msg := colorize.NewWriter(nil)
	msg.WriteString("[")
	msg.Fore = colorize.RED
	msg.AddAttr(colorize.Blod)
	msg.WriteString("hello world")
	msg.ClearAttrs()
	msg.Fore = colorize.DEFAULT
	msg.WriteString("]\n")
}

func writeToFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal("open file error")
	}
	defer file.Close()

	msg := colorize.NewWriter(file)
	msg.Fore = colorize.GREEN
	msg.Back = colorize.WHITE
	msg.AddAttr(colorize.Blinking)
	msg.WriteString("hello")
	msg.Escape = true
	msg.WriteString(" escaped\n")
}

func main() {
	writeToStdout()
	writeToFile()
}
