package colorize

import "testing"

func Test_Format(t *testing.T) {
	msg := NewMessage()
	msg.Font = RED
	msg.WriteString("hello\n")
}

func Test_AddSetting(t *testing.T) {
	ss := NewMessage()
	ss.AddAttr(Blod)
	ss.WriteString("hello world\n")
}
