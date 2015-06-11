package colorize

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

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

type console_t int
type color_t int

const (
	BLACK color_t = iota
	RED
	GREEN
	YELLOW
	BLUE
	MAGENTA
	CYAN
	WHITE
	DEFAULT color_t = 9
	FORE    color_t = 30
	BACK    color_t = 40
)

const (
	Blod       console_t = 1
	Dim        console_t = 2
	Underlined console_t = 4
	Blinking   console_t = 5
	Reverse    console_t = 7
	Hidden     console_t = 8
)

const prefix = "\033["
const reset console_t = 20
const clear console_t = 0

type Writer struct {
	Attrs      []console_t
	ResetAttrs []console_t
	Fore       color_t
	Back       color_t
	AutoClear  bool
	Escape     bool
	IsHtml     bool
	Writer     io.Writer
}

func NewWriter(w interface{}) *Writer {
	wr := &Writer{
		Fore:   DEFAULT,
		Back:   DEFAULT,
		Writer: os.Stdout,
	}
	if w, ok := w.(io.Writer); ok {
		wr.Writer = w
	}

	if _, ok := w.(*os.File); ok {
		wr.Escape = true
	}
	return wr
}

func (w *Writer) AddAttr(s console_t) {
	w.Attrs = append(w.Attrs, s)
}

func (w *Writer) AddRestAttr(s console_t) {
	w.ResetAttrs = append(w.ResetAttrs, s)
}

func (w *Writer) ClearAttrs() {
	w.Attrs = w.Attrs[0:0]
	w.ResetAttrs = w.ResetAttrs[0:0]
}

func (w *Writer) formatFore() string {
	return fmt.Sprintf("%d", FORE+w.Fore)
}

func (w *Writer) formatBack() string {
	return fmt.Sprintf("%d", BACK+w.Back)
}

func (w *Writer) formatAttrs() string {
	format := ""
	for _, attr := range w.Attrs {
		format += fmt.Sprintf("%d,", attr)
	}

	return strings.TrimRight(format, ",")
}

func (w *Writer) left() string {
	left := prefix + w.formatFore() + ";" + w.formatBack() + ";" + w.formatAttrs()
	left = strings.TrimRight(left, ";") + "m"
	return left
}

func (w *Writer) right() string {
	return prefix + fmt.Sprintf("%d", clear) + "m"
}

func (w *Writer) WriteString(msg string) (int, error) {
	return w.Write([]byte(msg))
}

func (w *Writer) Write(msg []byte) (int, error) {
	if w.Escape {
		return w.Writer.Write(msg)
	}

	buf := w.left()
	_, err := w.Writer.Write([]byte(buf))
	if err != nil {
		return 0, errors.New("format error")
	}

	n, err := w.Writer.Write(msg)
	if err != nil {
		return n, err
	}

	buf = w.right()
	_, err = w.Writer.Write([]byte(buf))
	if err != nil {
		return n, errors.New("format error")
	}
	return n, nil
}

func (w *Writer) converToHtml(msg string) string {
	return ""
}
