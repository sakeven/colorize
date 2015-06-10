package main

import (
	"colorize"
	"log"
	"os"
)

func writeToStdout() {
	msg := colorize.NewMessage()
	msg.Font = colorize.RED
	msg.WriteString("hello world")
}

func writeToFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal("open file error")
	}
	defer file.Close()

	msg := colorize.DefaultWriter()
	msg.Writer = file
	msg.Font = colorize.GREEN
	msg.Background = colorize.WHITE
	msg.AddAttr(colorize.Blinking)
	msg.WriteString("hello")
	msg.Escape = true
	msg.WriteString("escaped")
}

func main() {
	writeToStdout()
	writeToFile()
}
