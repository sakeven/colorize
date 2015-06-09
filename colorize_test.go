package colorize

import (
	"log"
	"testing"
)

func Test_Format(t *testing.T) {
	msg := new(Message)
	msg.Message = "Hello"
	msg.Font = RED

	log.Println(msg.format())
}

func Test_AddSetting(t *testing.T) {
	ss := new(Message)
	ss.AddSetting(color_tp(1))
	log.Println(ss.Settings)
}
