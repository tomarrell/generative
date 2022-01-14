package main

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/tomarrell/graphics/examples"
	"github.com/tomarrell/graphics/frame"
)

const (
	canvasWidth  = 1280
	canvasHeight = 880
	margin       = 40
)

var frameColor = color.Black

func main() {
	dc := gg.NewContext(canvasWidth, canvasHeight)

	// rand.Seed(time.Now().Unix())

	bg := background(canvasWidth, canvasHeight)
	dc.DrawImage(bg.Image(), 0, 0)

	dc.Translate(margin, margin)

	// Left
	{
		const w, h, aw, ah, bw, sw = 400, 800, 200, 600, 45, 20
		fr := frame.Frame(w, h, bw, sw, frameColor)
		dc.Translate(-sw, 0)
		dc.DrawImage(fr, 0, 0)

		shells := examples.Shells(aw, ah)
		dc.Translate(center(w, aw)+sw, center(h, ah))
		dc.DrawImage(shells, 0, 0)
	}

	dc.Identity()
	dc.Translate(margin, margin)
	dc.Translate(400+margin, 0)

	// Top Right
	{
		const w, h, aw, ah, bw, sw = 800 - 40, 400, 600, 200, 45, 20
		fr := frame.Frame(w, h, bw, sw, frameColor)
		dc.Translate(-sw, 0)
		dc.DrawImage(fr, 0, 0)

		spl := examples.Splotches(aw, ah)
		dc.Translate(center(w, aw)+sw, center(h, ah))
		dc.DrawImage(spl, 0, 0)
	}

	dc.Identity()
	dc.Translate(margin, margin)
	dc.Translate(360+2*margin, 360+2*margin)

	// Bottom middle
	{
		const w, h, aw, ah, bw, sw = 360, 360, 200, 200, 45, 20
		fr := frame.Frame(w, h, bw, sw, frameColor)
		dc.Translate(-sw, 0)
		dc.DrawImage(fr, 0, 0)

		spl := examples.Splotches(aw, ah)
		dc.Translate(center(w, aw)+sw, center(h, ah))
		dc.DrawImage(spl, 0, 0)
	}

	dc.Identity()
	dc.Translate(2*margin, margin)
	dc.Translate(2*(360+margin), 360+2*margin)

	// Bottom right
	{
		const w, h, aw, ah, bw, sw = 360, 360, 200, 200, 45, 20
		fr := frame.Frame(w, h, bw, sw, frameColor)
		dc.Translate(-sw, 0)
		dc.DrawImage(fr, 0, 0)

		spl := examples.Splotches(aw, ah)
		dc.Translate(center(w, aw)+sw, center(h, ah))
		dc.DrawImage(spl, 0, 0)
	}

	dc.SavePNG("out.png")
}

func center(parent, child float64) float64 {
	return parent/2 - child/2
}

func background(w, h float64) *gg.Context {
	c := gg.NewContext(int(w), int(h))

	c.SetColor(color.RGBA{120, 120, 120, 255})
	c.DrawRectangle(0, 0, w, h)
	c.Fill()

	return c
}
