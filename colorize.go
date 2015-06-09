package colorize

import "fmt"

/*
\033[0m 关闭所有属性
\033[1m 设置高亮度
\033[4m 下划线
\033[5m 闪烁
\033[7m 反显
\033[8m 消隐
\033[30m 至 \33[37m 设置前景色
\033[40m 至 \33[47m 设置背景色
\033[nA 光标上移n行
\033[nB 光标下移n行
\033[nC 光标右移n行
\033[nD 光标左移n行
\033[y;xH设置光标位置
\033[2J 清屏
\033[K 清除从光标到行尾的内容
\033[s 保存光标位置
\033[u 恢复光标位置
\033[?25l 隐藏光标
\033[?25h 显示光标
*/

type color_tp int

const (
	BLACK color_tp = iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
)

const (
	Dim        color_tp = 2
	Underlined color_tp = 4
	Blinking   color_tp = 5
	Reverse    color_tp = 7
	Hidden     color_tp = 8
)

const prefix = "\033["
const reset color_tp = 20
const clear = 0
const FONT color_tp = 30
const BACKGROUND color_tp = 40

func (c *color_tp) format(_type color_tp) string {
	gap := 30
	switch _type {
	case FONT:
		gap = 30
	case BACKGROUND:
		gap = 40
	default:
		panic("type error")
	}
	return fmt.Sprintf("%d", color_tp(gap)+*c)
}

type Message struct {
	Settings   []color_tp
	Font       color_tp
	Background color_tp
	Message    string
}

func (m *Message) AddSetting(s color_tp) {
	m.Settings = append(m.Settings, s)
}

func (m *Message) format() string {
	return prefix + m.Font.format(FONT) + "m" + m.Message + prefix + fmt.Sprintf("%d", clear) + "m"
}
