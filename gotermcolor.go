package gotermcolor

import (
	"bytes"
	"fmt"
)

type ColorString struct {
	StrAttr
	rawStr string
}

type StrAttr struct {
	attr      string
	forecolor string
	backcolor string
}

const STR_BEGIN = "\033["
const STR_END = "\033[0m"

type Mode string

const (
	RESET      = "0" //关闭所有属性
	BRIGHT     = "1" // 高亮度
	DIM        = "2"
	UNDER_LINE = "4" // 下划线
	BLINK      = "5" // 闪烁
	REVERSE    = "7" // 反显示
	HIDDEN     = "8" //消隐
)

type Color int

const (
	BLACK = iota
	RED
	GREEN
	YELLOW
	Blue
	MEGENTA
	CYAN
	WHITE
)

func (this *StrAttr) SetMode(attr Mode) {
	this.attr = string(attr)
}

func (this *StrAttr) SetForeColorRGB(r, g, b int) {
	this.forecolor = fmt.Sprintf("38;2;%d;%d;%d", r, g, b)
}

func (this *StrAttr) SetForeColor(c Color) {
	this.forecolor = fmt.Sprintf("3%v", c)
}

func (this *StrAttr) SetBackColorRGB(r, g, b int) {
	this.backcolor = fmt.Sprintf("48;2;%d;%d;%d", r, g, b)
}

func (this *StrAttr) SetBackColor(c Color) {
	this.backcolor = fmt.Sprintf("4%d", c)
}

func (this *StrAttr) ToString() string {
	buf := bytes.NewBufferString(this.attr)
	if this.forecolor != "" {
		if buf.String() != "" {
			buf.WriteString(";")
		}
		buf.WriteString(this.forecolor)
	}

	if this.backcolor != "" {
		if buf.String() != "" {
			buf.WriteString(";")
		}
		buf.WriteString(this.backcolor)
	}

	buf.WriteString("m")
	return STR_BEGIN + buf.String()
}

func NewColorString(raw string) ColorString {
	return ColorString{
		rawStr: raw,
	}
}
func (this *ColorString) ToString() string {
	return this.StrAttr.ToString() + this.rawStr + STR_END
}

func (this *ColorString) GetStrAttr() StrAttr {
	return this.StrAttr
}

func (this *ColorString) SetStrAttr(s StrAttr) {
	this.StrAttr = s
}
