package main

import (
	"image/color"

	"github.com/fogleman/gg"
)

const (
	canvasWidth  = 1300
	canvasHeight = 1000

	frameWidth       = 900
	frameHeight      = 800
	frameBorderWidth = 45
	shadowWidth      = 20

	artWidth  = frameWidth / 2
	artHeight = frameHeight / 2
)

var red = color.RGBA{255, 0, 0, 255}

func main() {
	dc := gg.NewContext(canvasWidth, canvasHeight)

	bg := background(canvasWidth, canvasHeight)
	dc.DrawImage(bg.Image(), 0, 0)

	fr := frame(frameWidth, frameHeight, frameBorderWidth)
	dc.Translate((canvasWidth-frameWidth)/2, (canvasHeight-frameHeight)/2)
	dc.DrawImage(fr.Image(), 0, 0)

	a := art(artWidth, artHeight)
	dc.Translate(artWidth/2+shadowWidth, artHeight/2)
	dc.DrawImage(a.Image(), 0, 0)

	dc.SavePNG("out.png")
}

func art(w, h float64) *gg.Context {
	c := gg.NewContext(int(w), int(h))

	c.SetColor(color.White)
	c.DrawRectangle(0, 0, w, h)
	c.Fill()

	return c
}

func background(w, h float64) *gg.Context {
	c := gg.NewContext(int(w), int(h))

	c.SetColor(color.RGBA{100, 100, 100, 255})
	c.DrawRectangle(0, 0, w, h)
	c.Fill()

	return c
}

func frame(w, h, bw float64) *gg.Context {
	c := gg.NewContext(int(w)+shadowWidth, int(h)+shadowWidth)
	c.InvertY()
	c.RotateAbout(gg.Radians(180), float64(int(w)+shadowWidth)/2, float64(int(h)+shadowWidth)/2)

	c.SetColor(color.RGBA{0, 0, 0, 255})
	c.DrawRectangle(0, 0, w, h)
	c.Fill()

	runs := 20
	for i := runs; i > 0; i-- {
		c.SetColor(color.RGBA{0, 0, 0, uint8(runs - i)})
		c.DrawRectangle(float64(i*1), float64(i*1), w, h)
		c.Fill()
	}

	c.SetColor(color.RGBA{245, 245, 245, 255})
	c.DrawRectangle(bw, bw, w-2*bw, h-2*bw)
	c.Fill()

	return c
}