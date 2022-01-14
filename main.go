package main

import (
	"image/color"
	"math"

	"github.com/fogleman/gg"
	"github.com/tomarrell/graphics/examples"
	"github.com/tomarrell/graphics/frame"
)

const (
	canvasWidth  = 1300
	canvasHeight = 1000

	frameWidth       = canvasWidth / 2
	frameHeight      = canvasHeight / 2
	frameBorderWidth = 45
	frameShadowWidth = 20
)

var (
	artWidth  = frameWidth - math.Min(frameWidth/2, frameHeight/2)
	artHeight = frameHeight - math.Min(frameWidth/2, frameHeight/2)
)

var red = color.RGBA{255, 0, 0, 255}

func main() {
	dc := gg.NewContext(canvasWidth, canvasHeight)

	bg := background(canvasWidth, canvasHeight)
	dc.DrawImage(bg.Image(), 0, 0)

	fr := frame.Frame(frameWidth, frameHeight, frameBorderWidth, frameShadowWidth)
	dc.DrawImage(fr, 0, 0)

	a := examples.Shells(artWidth, artHeight)
	dc.Translate((frameWidth/2-artWidth/2)+frameShadowWidth, (frameHeight/2 - artHeight/2))
	dc.DrawImage(a, 0, 0)

	dc.SavePNG("out.png")
}

func background(w, h float64) *gg.Context {
	c := gg.NewContext(int(w), int(h))

	c.SetColor(color.RGBA{100, 100, 100, 255})
	c.DrawRectangle(0, 0, w, h)
	c.Fill()

	return c
}
