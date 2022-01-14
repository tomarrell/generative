package main

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
)

const (
	canvasWidth  = 1300
	canvasHeight = 1000

	frameWidth       = 700
	frameHeight      = 800
	frameBorderWidth = 45
	shadowWidth      = 20
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

	fr := frame(frameWidth, frameHeight, frameBorderWidth)
	dc.Translate((canvasWidth-frameWidth)/2, (canvasHeight-frameHeight)/2)
	dc.DrawImage(fr.Image(), 0, 0)

	a := art(artWidth, artHeight)
	dc.Translate((frameWidth/2-artWidth/2)+shadowWidth, (frameHeight/2 - artHeight/2))
	dc.DrawImage(a.Image(), 0, 0)

	dc.SavePNG("out.png")
}

func art(w, h float64) *gg.Context {
	c := gg.NewContext(int(w), int(h))
	insetbg(c, w, h)

	c.Scale(0.1, 0.1)

	for i := 0; i < 100; i++ {
		for j := 0; j < 10; j++ {
			shell(c, w, h)
			c.Translate(w, 0)
		}
		c.Translate(-10*w, h)
	}

	return c
}

func shell(c *gg.Context, w, h float64) {
	n := gg.NewContext(int(w), int(h))

	n.SetLineWidth(3 * 5)
	n.SetColor(color.Black)

	n.MoveTo(w/2, h/2)
	n.LineTo(w/4, h/2)
	n.LineTo(w/4, h/4)
	n.LineTo(w*3/4, h/4)
	n.LineTo(w*3/4, h*3/4)
	n.LineTo(w/4, h*3/4)

	n.Stroke()

	var img image.Image
	if rand.Intn(2) > 0 {
		img = imaging.FlipV(n.Image())
	} else {
		img = n.Image()
	}

	if rand.Intn(2) > 0 {
		img = imaging.FlipH(img)
	}

	c.DrawImage(img, 0, 0)
}

func insetbg(c *gg.Context, w, h float64) {
	c.SetColor(color.RGBA{0, 0, 0, 10})
	c.DrawRectangle(0, 0, w, h)
	c.Fill()

	c.Translate(0, 2)
	c.SetColor(color.White)
	c.DrawRectangle(0, 0, w-2, h)
	c.Fill()

	c.Identity()
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
