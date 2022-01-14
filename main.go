package main

import (
	"image/color"
	"math"

	"github.com/fogleman/gg"
	"github.com/tomarrell/graphics/examples"
	"github.com/tomarrell/graphics/frame"
)

const (
	canvasWidth  = 1300 + 2*frameShadowWidth
	canvasHeight = 1000 + frameShadowWidth

	frameWidth       = canvasWidth/2 - frameShadowWidth
	frameHeight      = canvasHeight/2 - frameShadowWidth/2
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

	// Shells
	{
		fr := frame.Frame(frameWidth, 2*frameHeight, frameBorderWidth, frameShadowWidth)
		dc.DrawImage(fr, 0, 0)

		shells := examples.Shells(artWidth, 2*artHeight)
		dc.Translate((frameWidth/2-artWidth/2)+frameShadowWidth, (frameHeight - artHeight))
		dc.DrawImage(shells, 0, 0)
	}

	// Dots
	{
		dc.Identity()
		dc.Translate(frameWidth+frameShadowWidth, 0)

		fr := frame.Frame(frameWidth, frameHeight, frameBorderWidth, frameShadowWidth)
		dc.DrawImage(fr, 0, 0)

		spl := examples.Splotches(artWidth, artHeight)
		dc.Translate((frameWidth/2-artWidth/2)+frameShadowWidth, (frameHeight/2 - artHeight/2))
		dc.DrawImage(spl, 0, 0)
	}

	dc.SavePNG("out.png")
}

func background(w, h float64) *gg.Context {
	c := gg.NewContext(int(w), int(h))

	// c.SetColor(color.RGBA{140, 120, 140, 255})
	// c.DrawRectangle(0, 0, w, h)
	// c.Fill()

	return c
}
