package main

import (
	"image/color"

	"github.com/fogleman/gg"
)

const (
	canvasWidth  = 1300
	canvasHeight = 1000

	frameWidth       = 700
	frameHeight      = 800
	frameBorderWidth = 50
)

func main() {
	dc := gg.NewContext(canvasWidth, canvasHeight)

	bg := background(canvasWidth, canvasHeight)
	dc.DrawImage(bg.Image(), 0, 0)

	fr := frame(frameWidth, frameHeight, frameBorderWidth)
	dc.Translate((canvasWidth-frameWidth)/2, (canvasHeight-frameHeight)/2)
	dc.DrawImage(fr.Image(), 0, 0)

	fr2 := frame(frameWidth, frameHeight, frameBorderWidth)
	dc.Translate(-10, 10)
	fr2.SetColor(color.RGBA{0, 0, 0, 100})
	fr2.Fill()
	dc.DrawImage(fr2.Image(), 0, 0)

	dc.SavePNG("out.png")
}

func background(w, h float64) *gg.Context {
	c := gg.NewContext(int(w), int(h))

	c.SetColor(color.White)
	c.DrawRectangle(0, 0, w, h)
	c.Fill()

	return c
}

func frame(w, h, bw float64) *gg.Context {
	c := gg.NewContext(int(w), int(h))

	c.SetColor(color.Black)
	c.DrawRectangle(0, 0, w, h)
	c.Fill()

	c.SetColor(color.White)
	c.DrawRectangle(bw, bw, w-2*bw, h-2*bw)
	c.Fill()

	return c
}
