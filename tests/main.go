package main

import (
	"fmt"
	gtc "gotermcolor"
)

func main() {
	h := gtc.NewColorString("hello world")
	h.SetMode(gtc.DIM)
	h.SetBackColor(gtc.BLACK)
	h.SetForeColor(gtc.GREEN)
	fmt.Println(h.ToString())

	b := gtc.NewColorString("huaer zheyanghong")
	b.SetStrAttr(h.GetStrAttr())
	b.SetForeColorRGB(255, 255, 255)
	fmt.Println(b.ToString())
	c := gtc.NewColorString("hello world")
	bb := c.SetForeColorRGB(255, 255, 255).SetMode(gtc.UNDER_LINE).ToString()
	fmt.Println(bb)
	fmt.Println(gtc.NewColorString("hello world").SetForeColor(gtc.RED).SetBackColor(gtc.GREEN).ToString())
}
