package examples

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/tomarrell/graphics/frame"
)

func Shells(width, height float64) image.Image {
	c := gg.NewContext(int(width), int(height))

	c.DrawImage(frame.Artboard(width, height), 0, 0)

	c.Scale(0.1, 0.1)

	for i := 0; i < 100; i++ {
		for j := 0; j < 10; j++ {
			shell(c, width, height)
			c.Translate(width, 0)
		}
		c.Translate(-10*width, height)
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
