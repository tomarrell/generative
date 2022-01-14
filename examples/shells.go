package examples

import (
	"image"
	"image/color"
	"math"
	"math/rand"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/tomarrell/graphics/frame"
)

const scale = 0.1

func Shells(width, height float64) image.Image {
	c := gg.NewContext(int(width), int(height))

	c.DrawImage(frame.Artboard(width, height), 0, 0)
	c.Scale(scale, scale)

	min := math.Min(width, height)

	for i := 0; i < int(height/min*1/scale); i++ {
		c.Push()
		for j := 0; j < 10; j++ {
			shell(c, min, min)
			c.Translate(min, 0)
		}
		c.Pop()
		c.Translate(0, min)
	}

	return c.Image()
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
