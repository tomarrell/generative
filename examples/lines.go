package examples

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
	"github.com/tomarrell/graphics/frame"
)

func Lines(width, height float64) image.Image {
	c := gg.NewContext(int(width), int(height))
	c.DrawImage(frame.Artboard(width, height), 0, 0)

	c.Translate(0, 5)
	c.SetLineWidth(2)

	for i := 0; i < 10; i++ {
		y := float64(10 * i)
		c.SetColor(color.Black)
		c.DrawLine(0, y, width, y)
		c.Stroke()
	}

	c.Translate(0, height/2-20)
	c.SetLineWidth(2)

	for i := 0; i < 20; i++ {
		y := float64(10 * i)
		c.SetColor(color.Black)
		c.DrawLine(0, y, width, y)
		c.Stroke()
	}

	return c.Image()
}
